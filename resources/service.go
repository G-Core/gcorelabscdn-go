package resources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/skyeng/gcorelabscdn-go/gcore"
)

var _ ResourceService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodPost, "/cdn/resources", req, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/cdn/resources/%d", id), nil, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}

func (s *Service) Update(ctx context.Context, id int64, req *UpdateRequest) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodPut, fmt.Sprintf("/cdn/resources/%d", id), req, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}
