package origingroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type OriginGroupService interface {
	Create(ctx context.Context, req *GroupRequest) (*OriginGroup, error)
	Get(ctx context.Context, id int64) (*OriginGroup, error)
	Update(ctx context.Context, id int64, req *GroupRequest) (*OriginGroup, error)
	Delete(ctx context.Context, id int64) error
}

var _ OriginGroupService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, req *GroupRequest) (*OriginGroup, error) {
	var group OriginGroup
	if err := s.r.Request(ctx, http.MethodPost, "/originGroups", req, &group); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &group, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*OriginGroup, error) {
	var group OriginGroup
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/originGroups/%d", id), nil, &group); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &group, nil
}

func (s *Service) Update(ctx context.Context, id int64, req *GroupRequest) (*OriginGroup, error) {
	var group OriginGroup
	if err := s.r.Request(ctx, http.MethodPut, fmt.Sprintf("/originGroups/%d", id), req, &group); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &group, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if err := s.r.Request(ctx, http.MethodDelete, fmt.Sprintf("/originGroups/%d", id), nil, nil); err != nil {

		return fmt.Errorf("request: %w", err)
	}

	return nil
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
