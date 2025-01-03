package cacerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*CACert, error) {
	var cert CACert

	if err := s.r.Request(ctx, http.MethodPost, "/cdn/sslCertificates", req, &cert); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &cert, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*CACert, error) {
	var cert CACert
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/cdn/sslCertificates/%d", id), nil, &cert); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &cert, nil
}

func (s *Service) Update(ctx context.Context, id int64, req *UpdateRequest) (*CACert, error) {
	var cert CACert
	if err := s.r.Request(ctx, http.MethodPatch, fmt.Sprintf("/cdn/sslCertificates/%d", id), req, &cert); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &cert, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if err := s.r.Request(ctx, http.MethodDelete, fmt.Sprintf("/cdn/sslCertificates/%d", id), nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
