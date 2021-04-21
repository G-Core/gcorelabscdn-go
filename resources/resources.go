package resources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
	"github.com/G-Core/gcorelabscdn-go/resources/rules"
)

type ResourceService interface {
	Rules(opts ...rules.ServiceOption) rules.RulesService
	Create(ctx context.Context, req *CreateRequest) (*Resource, error)
}

type Service struct {
	root gcore.Pather
	r    gcore.Requester
	path string
}

type ServiceOption func(*Service)

func WithID(id int64) ServiceOption {
	return func(s *Service) {
		s.path = fmt.Sprintf("%s/%d", s.path, id)
	}
}

func NewService(root gcore.Pather, r gcore.Requester, opts ...ServiceOption) *Service {
	s := &Service{root: root, r: r, path: "/resources"}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Service) Path() string {
	return s.root.Path() + s.path
}

func (s *Service) Rules(opts ...rules.ServiceOption) rules.RulesService {
	return rules.NewService(s, s.r, opts...)
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodPost, s.path, req, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}

type CreateRequest struct {
	Cname              string
	OriginGroup        int
	Origin             string
	SecondaryHostnames []string
}

type Resource struct {
	ID                 int64
	Status             string
	Client             int64
	OriginGroup        int64
	Cname              string
	SecondaryHostnames []string
	Shielded           bool
}
