package rules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type RulesService interface {
	Create(ctx context.Context, req *CreateRequest) (*Rule, error)
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
	s := &Service{root: root, r: r, path: "/rules"}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Service) Path() string {
	return s.root.Path() + s.path
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*Rule, error) {
	var rule Rule
	if err := s.r.Request(ctx, http.MethodPost, s.Path(), req, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

type CreateRequest struct {
	Name     string
	Rule     string
	RuleType int
}

type Rule struct {
	ID      int64
	Name    string
	Deleted bool
	Pattern string `json:"rule"`
	Type    int    `json:"ruleType"`
}
