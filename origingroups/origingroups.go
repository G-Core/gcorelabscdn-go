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
	Name    string          `json:"name"`
	UseNext bool            `json:"useNext"`
	Origins []OriginRequest `json:"origins"`
}

type OriginRequest struct {
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}

type OriginGroup struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	UseNext bool     `json:"useNext"`
	Origins []Origin `json:"origin_ids"`
}

type Origin struct {
	ID      int64  `json:"id"`
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}
