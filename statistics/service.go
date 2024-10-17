// Declare the API service for the statistics endpoint https://api.gcore.com/docs/cdn#tag/Statistics
package statistics

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ StatisticsService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) GetAggregatedStatistics(ctx context.Context, req *GetAggregatedStatisticsRequest) (*AggregatedResource, error) {
	var aggregatedresource AggregatedResource
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/aggregate/stats", nil, &aggregatedresource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &aggregatedresource, nil
}

func (s *Service) GetOriginShieldingUsageStatistics(ctx context.Context, req *GetOriginShieldingUsageStatisticsRequest) (*OriginShieldingUsage, error) {
	var originshieldingusage OriginShieldingUsage
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/shield_usage/series", nil, &originshieldingusage); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &originshieldingusage, nil
}

func (s *Service) GetAggregatedOriginShieldingUsageStatistics(ctx context.Context, req *GetAggregatedOriginShieldingUsageStatisticsRequest) (*AggregatedOriginShieldingUsage, error) {
	var aggregatedoriginshieldingusage AggregatedOriginShieldingUsage
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/shield_usage/aggregated", nil, &aggregatedoriginshieldingusage); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &aggregatedoriginshieldingusage, nil
}

func (s *Service) GetRawLogsUsageStatistics(ctx context.Context, req *GetRawLogsUsageStatisticsRequest) (*OriginShieldingUsage, error) {
	var originshieldingusage OriginShieldingUsage
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/raw_logs_usage/series", nil, &originshieldingusage); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &originshieldingusage, nil
}

func (s *Service) GetAggregatedRawLogsUsageStatistics(ctx context.Context, req *GetAggregatedRawLogsUsageStatisticsRequest) (*AggregatedRowLogUsage, error) {
	var aggregatedrowlogusage AggregatedRowLogUsage
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/raw_logs_usage/aggregated", nil, &aggregatedrowlogusage); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &aggregatedrowlogusage, nil
}

func (s *Service) GetNetworkCapacity(ctx context.Context) (*NetworkCapacity, error) {
	var networkcapacity NetworkCapacity
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/advanced/v1/capacity", nil, &networkcapacity); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &networkcapacity, nil
}

func (s *Service) GetCDNResourceStatistics(ctx context.Context, req *GetCDNMetricsRequest) (string, error) {
	request, err := buildRequest(ctx, req)
	if err != nil {
		return "", err
	}

	// Perform request.
	client := http.Client{Timeout: time.Minute}
	res, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := readBodyAndCheckStatus(res)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// buildRequest creates a *http.Request for the GetCDNMetrics endpoint, adds all query parameters and headers and returns it.
func buildRequest(ctx context.Context, req *GetCDNMetricsRequest) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.gcore.com/cdn/statistics/series", http.NoBody)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("service", "CDN")
	q.Add("from", req.From)
	q.Add("to", req.To)
	q.Add("granularity", req.Granularity)

	for _, metric := range req.Metrics {
		q.Add("metrics", metric)
	}

	for _, group := range req.GroupBy {
		q.Add("group_by", group)
	}
	request.URL.RawQuery = q.Encode()

	request.Header.Add("Authorization", fmt.Sprintf("APIKey %s", req.ApiKey))
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

func readBodyAndCheckStatus(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch data from Gcore: %s. status code: %d", string(body), res.StatusCode)
	}

	return body, nil
}
