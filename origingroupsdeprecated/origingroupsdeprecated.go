package origingroupsdeprecated

import (
	"context"
)

type OriginGroupServiceDeprecated interface {
	Create(ctx context.Context, req *GroupRequestDeprecated) (*OriginGroupDeprecated, error)
	Get(ctx context.Context, id int64) (*OriginGroupDeprecated, error)
	Update(ctx context.Context, id int64, req *GroupRequestDeprecated) (*OriginGroupDeprecated, error)
	Delete(ctx context.Context, id int64) error
}

type GroupRequestDeprecated struct {
	Name    string          `json:"name"`
	UseNext bool            `json:"useNext"`
	Origins []OriginRequestDeprecated `json:"origins"`
}

type OriginRequestDeprecated struct {
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}

type OriginGroupDeprecated struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	UseNext bool     `json:"useNext"`
	Origins []OriginDeprecated `json:"origin_ids"`
}

type OriginDeprecated struct {
	ID      int64  `json:"id"`
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}
