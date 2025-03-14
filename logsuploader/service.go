package logsuploader

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ LogsUploaderService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) PolicyCreate(ctx context.Context, req *PolicyCreateRequest) (*Policy, error) {
	var policy Policy

	if err := s.r.Request(ctx, http.MethodPost, "/cdn/logs_uploader/policies", req, &policy); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &policy, nil
}

func (s *Service) PolicyList(ctx context.Context, limit, offset int) ([]Policy, error) {
	var policies []Policy
	params := url.Values{}
	if limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", limit))
	}
	if offset > 0 {
		params.Add("offset", fmt.Sprintf("%d", offset))
	}

	path := "/cdn/logs_uploader/policies"
	if len(params) > 0 {
		path = fmt.Sprintf("%s?%s", path, params.Encode())
	}

	if err := s.r.Request(ctx, http.MethodGet, path, nil, &policies); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return policies, nil
}

func (s *Service) PolicyGet(ctx context.Context, id int64) (*Policy, error) {
	var policy Policy
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/cdn/logs_uploader/policies/%d", id), nil, &policy); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &policy, nil
}

func (s *Service) PolicyUpdate(ctx context.Context, id int64, req *PolicyUpdateRequest) (*Policy, error) {
	var policy Policy
	if err := s.r.Request(ctx, http.MethodPatch, fmt.Sprintf("/cdn/logs_uploader/policies/%d", id), req, &policy); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &policy, nil
}

func (s *Service) PolicyDelete(ctx context.Context, id int64) error {
	if err := s.r.Request(ctx, http.MethodDelete, fmt.Sprintf("/cdn/logs_uploader/policies/%d", id), nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
