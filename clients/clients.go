package clients

import (
	"context"
	"time"
)

type ClientsMeService interface {
    Get(ctx context.Context) (*ClientsMe, error)
    Update(ctx context.Context, req *ClientsMeUpdateRequest) (*ClientsMe, error)
}

type ServiceInfo struct {
	Enabled bool      `json:"enabled"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated,omitempty"`
}

type ClientsMe struct {
	ID                        int64        `json:"id"`
	Cname                     string       `json:"cname"`
	Created                   time.Time    `json:"created,omitempty"`
	Updated                   time.Time    `json:"updated,omitempty"`
	UtilizationLevel          int          `json:"utilization_level"`
	UseBalancer               bool         `json:"use_balancer"`
	AutoSuspendEnabled        bool         `json:"auto_suspend_enabled"`
	CDNResourcesRulesMaxCount int          `json:"cdn_resources_rules_max_count,omitempty"`
	Service                   *ServiceInfo `json:"service,omitempty"`
}


type ClientsMeUpdateRequest struct {
    UtilizationLevel int `json:"utilization_level"`
}
