package resources

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcorelabscdn-go/gcore"
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

func (s *Service) List(ctx context.Context, limit, offset int) ([]Resource, error) {
	var resources []Resource
	params := url.Values{}
	if limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}
	if offset > 0 {
		params.Add("offset", fmt.Sprintf("%d", offset))
	}

	path := "/cdn/resources"
	if len(params) > 0 {
		path = fmt.Sprintf("%s?%s", path, params.Encode())
	}

	if err := s.r.Request(ctx, http.MethodGet, path, nil, &resources); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return resources, nil
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

func (s *Service) Delete(ctx context.Context, resourceID int64) error {
	path := fmt.Sprintf("/cdn/resources/%d", resourceID)

	var body ActivateRequest = ActivateRequest{
		Active: false,
	}
	var resource Resource

	// deactivate the resource instance before deletion
	if err := s.r.Request(ctx, http.MethodPatch, path, &body, &resource); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	if err := s.r.Request(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
