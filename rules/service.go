package rules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ RulesService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, resourceID int64, req *CreateRequest) (*Rule, error) {
	var rule Rule

	path := fmt.Sprintf("/cdn/resources/%d/rules", resourceID)
	if err := s.r.Request(ctx, http.MethodPost, path, req, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

func (s *Service) Get(ctx context.Context, resourceID, ruleID int64) (*Rule, error) {
	var rule Rule

	path := fmt.Sprintf("/cdn/resources/%d/rules/%d", resourceID, ruleID)
	if err := s.r.Request(ctx, http.MethodGet, path, nil, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

func (s *Service) Update(ctx context.Context, resourceID, ruleID int64, req *UpdateRequest) (*Rule, error) {
	var rule Rule

	path := fmt.Sprintf("/cdn/resources/%d/rules/%d", resourceID, ruleID)
	if err := s.r.Request(ctx, http.MethodPut, path, req, &rule); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &rule, nil
}

func (s *Service) Delete(ctx context.Context, resourceID, ruleID int64) error {
	path := fmt.Sprintf("/cdn/resources/%d/rules/%d", resourceID, ruleID)

	var body ActivateRequest = ActivateRequest{
		Active: false,
	}
	var rule Rule

	// deactivate the rule instance before deletion
	if err := s.r.Request(ctx, http.MethodPatch, path, &body, &rule); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	if err := s.r.Request(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
