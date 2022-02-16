package origingroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ OriginGroupService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, req *GroupRequest) (*OriginGroup, error) {
	var group OriginGroup
	if err := s.r.Request(ctx, http.MethodPost, "/cdn/originGroups", req, &group); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &group, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*OriginGroup, error) {
	var group OriginGroup
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/cdn/originGroups/%d", id), nil, &group); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &group, nil
}

func (s *Service) Update(ctx context.Context, id int64, req *GroupRequest) (*OriginGroup, error) {
	var group OriginGroup
	if err := s.r.Request(ctx, http.MethodPut, fmt.Sprintf("/cdn/originGroups/%d", id), req, &group); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &group, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if err := s.r.Request(ctx, http.MethodDelete, fmt.Sprintf("/cdn/originGroups/%d", id), nil, nil); err != nil {

		return fmt.Errorf("request: %w", err)
	}

	return nil
}
