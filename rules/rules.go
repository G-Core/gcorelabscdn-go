package rules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type RulesService interface {
	Create(ctx context.Context, resourceID int64, req *CreateRequest) (*Rule, error)
	Get(ctx context.Context, resourceID, ruleID int64) (*Rule, error)
	Update(ctx context.Context, resourceID, ruleID int64, req *UpdateRequest) (*Rule, error)
	Delete(ctx context.Context, resourceID, ruleID int64) error
}

var _ RulesService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, resourceID int64, req *CreateRequest) (*Rule, error) {
	var rule Rule

	path := fmt.Sprintf("/resources/%d/rules", resourceID)
	if err := s.r.Request(ctx, http.MethodPost, path, req, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

func (s *Service) Get(ctx context.Context, resourceID, ruleID int64) (*Rule, error) {
	var rule Rule

	path := fmt.Sprintf("/resources/%d/rules/%d", resourceID, ruleID)
	if err := s.r.Request(ctx, http.MethodGet, path, nil, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

func (s *Service) Update(ctx context.Context, resourceID, ruleID int64, req *UpdateRequest) (*Rule, error) {
	var rule Rule

	path := fmt.Sprintf("/resources/%d/rules/%d", resourceID, ruleID)
	if err := s.r.Request(ctx, http.MethodPut, path, req, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

func (s *Service) Delete(ctx context.Context, resourceID, ruleID int64) error {
	path := fmt.Sprintf("/resources/%d/rules/%d", resourceID, ruleID)
	if err := s.r.Request(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}

type CreateRequest struct {
	Name     string `json:"name"`
	Rule     string `json:"rule"`
	RuleType int    `json:"ruleType"`
}

type UpdateRequest struct {
	Name     string `json:"name"`
	Rule     string `json:"rule"`
	RuleType int    `json:"ruleType"`
}

type Rule struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Deleted     bool   `json:"deleted"`
	OriginGroup int    `json:"originGroup"`
	Pattern     string `json:"rule"`
	Type        int    `json:"ruleType"`
}
