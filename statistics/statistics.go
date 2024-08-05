// Declare the schemas for the statistics endpoint https://api.gcore.com/docs/cdn#tag/Statistics
// the Conversion from the OpenAPI schema to the Go struct is done by https://transform.tools/json-to-go
package statistics

import (
	"context"
	"time"
)

type StatisticsService interface {
	GetAggregatedStatistics(ctx context.Context, req *GetAggregatedStatisticsRequest) (*AggregatedResource, error)
	GetOriginShieldingUsageStatistics(ctx context.Context, req *GetOriginShieldingUsageStatisticsRequest) (*OriginShieldingUsage, error)
	GetAggregatedOriginShieldingUsageStatistics(ctx context.Context, req *GetAggregatedOriginShieldingUsageStatisticsRequest) (*AggregatedOriginShieldingUsage, error)
	GetRawLogsUsageStatistics(ctx context.Context, req *GetRawLogsUsageStatisticsRequest) (*OriginShieldingUsage, error)
	GetAggregatedRawLogsUsageStatistics(ctx context.Context, req *GetAggregatedRawLogsUsageStatisticsRequest) (*AggregatedRowLogUsage, error)
	GetNetworkCapacity(ctx context.Context) (*NetworkCapacity, error)
	GetCDNResourceStatistics(ctx context.Context, req *GetCDNMetricsRequest) (string, error)
}

type GetCDNResourceStatisticsRequest struct {
	Service     string `json:"service"`
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity"`
	Metrics     string `json:"metrics"`
	GroupBy     string `json:"group_by"`
	Countries   string `json:"countries"`
	Regions     string `json:"regions"`
	Resource    int    `json:"resource"`
}

type GetAggregatedStatisticsRequest struct {
	Service   string `json:"service"`
	From      string `json:"from"`
	To        string `json:"to"`
	Metrics   string `json:"metrics"`
	GroupBy   string `json:"group_by"`
	Regions   string `json:"regions"`
	Countries string `json:"countries"`
	Resource  int    `json:"resource"`
	Flat      bool   `json:"flat"`
}

type GetOriginShieldingUsageStatisticsRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Resource int    `json:"resource"`
}

type GetAggregatedOriginShieldingUsageStatisticsRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	GroupBy  string `json:"group_by"`
	Resource int    `json:"resource"`
	Flat     bool   `json:"flat"`
}

type GetRawLogsUsageStatisticsRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Resource int    `json:"resource"`
}

type GetAggregatedRawLogsUsageStatisticsRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	GroupBy  string `json:"group_by"`
	Resource int    `json:"resource"`
	Flat     bool   `json:"flat"`
}

type FilterGroup struct {
	Field string   `json:"field"`
	Op    string   `json:"op"`
	Value []string `json:"value"` // array of string or int
}

type GetCDNMetricsRequest struct {
	ApiKey      string   `json:"-"`
	From        string   `json:"from"` // ISO 8601/RFC 3339 format, UTC.
	To          string   `json:"to"`   // ISO 8601/RFC 3339 format, UTC.
	Granularity string   `json:"granularity"`
	Metrics     []string `json:"metrics"`
	GroupBy     []string `json:"group_by"`
}

// Main responses schemas

type CDNResource struct {
	Resource struct {
		Num struct {
			Region Region `json:"region"`
		}
	} `json:"resource"`
}

type AggregatedResource struct {
	Resource struct {
		Num struct {
			Region AggregatedRegion `json:"region"`
		}
	} `json:"resource"`
}

type OriginShieldingUsage struct {
	ID         int       `json:"id"`
	ActiveFrom time.Time `json:"active_from"`
	ActiveTo   string    `json:"active_to"`
	ClientID   int       `json:"client_id"`
	ResourceID int       `json:"resource_id"`
	Cname      string    `json:"cname"`
}

type AggregatedOriginShieldingUsage struct { // nested response object
	Resource struct {
		Num1 struct {
			Metrics MetricsShieldUsage `json:"metrics"`
		} `json:"1"`
	} `json:"resource"`
}

type AggregatedRowLogUsage struct { // nested response object
	Metrics struct {
		RawLogsUsage int `json:"raw_logs_usage"`
	} `json:"metrics"`
}

type NetworkCapacity []struct {
	CountryCode string  `json:"country_code"`
	Country     string  `json:"country"`
	Capacity    float64 `json:"capacity"`
}

type CDNMetricsData struct {
	Data []CDNMetricsDataResponse `json:"data"`
}

type CDNMetricsDataResponse struct {
	EdgeStatus2Xx int `json:"edge_status_2xx,omitempty"`
	Timestamp     int `json:"timestamp,omitempty"`
}

type ResourceID struct {
	ID int `json:"id"`
}

type Region struct {
	Asia  Metrics `json:"asia"`
	Cis   Metrics `json:"cis"`
	EU    Metrics `json:"eu"`
	Latam Metrics `json:"latam"`
	Me    Metrics `json:"me"`
	Na    Metrics `json:"na"`
}

type AggregatedRegion struct {
	Cis Metrics `json:"cis"`
	EU  Metrics `json:"eu"`
}

type Metrics struct {
	SentBytes      [][]int `json:"sent_bytes"`
	TotalBytes     [][]int `json:"total_bytes"`
	BackblazeBytes [][]int `json:"backblaze_bytes"`
	UpstreamBytes  [][]int `json:"upstream_bytes"`
}

type MetricsShieldUsage struct {
	ShieldUsage int `json:"shield_usage"`
}

type MetricsRowLogUsage struct {
	RawLogUsage int `json:"raw_log_usage"`
}
