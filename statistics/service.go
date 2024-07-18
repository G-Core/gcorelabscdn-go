// Declare the API service for the statistics endpoint https://api.gcore.com/docs/cdn#tag/Statistics
package statistics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ StatisticsService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func (s *Service) GetCDNResourceStatistics(ctx context.Context, req *GetCDNResourceStatisticsRequest) (*CDNResource, error) {
	var cdnresource CDNResource
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/series", nil, &cdnresource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &cdnresource, nil

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

func (s *Service) CreateCDNMetrics(ctx context.Context, req *CreateCDNMetricsRequest) (*CDNMetricsData, error) {
	var cdnmetricsdata CDNMetricsData
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/statistics/series", nil, &cdnmetricsdata); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &cdnmetricsdata, nil

}