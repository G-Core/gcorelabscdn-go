package origingroups

import (
	"context"
)

type OriginGroupService interface {
	Create(ctx context.Context, req *GroupRequest) (*OriginGroup, error)
	Get(ctx context.Context, id int64) (*OriginGroup, error)
	Update(ctx context.Context, id int64, req *GroupRequest) (*OriginGroup, error)
	Delete(ctx context.Context, id int64) error
}

type GroupRequest struct {
	Name              string          `json:"name"`
	UseNext           bool            `json:"use_next"`
	Sources           []SourceRequest `json:"sources"`
	ProxyNextUpstream []string        `json:"proxy_next_upstream"`
}

type SourceRequest struct {
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}

type OriginGroup struct {
	ID                int64    `json:"id"`
	Name              string   `json:"name"`
	UseNext           bool     `json:"use_next"`
	Sources           []Source `json:"sources"`
	ProxyNextUpstream []string `json:"proxy_next_upstream"`
}

type Source struct {
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}
