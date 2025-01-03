package ruletemplates

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

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*RuleTemplate, error) {
	var ruleTemplate RuleTemplate

	if err := s.r.Request(ctx, http.MethodPost, "/cdn/resources/rule_templates", req, &ruleTemplate); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &ruleTemplate, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*RuleTemplate, error) {
	var ruleTemplate RuleTemplate
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/cdn/resources/rule_templates/%d", id), nil, &ruleTemplate); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &ruleTemplate, nil
}

func (s *Service) Update(ctx context.Context, id int64, req *UpdateRequest) (*RuleTemplate, error) {
	var ruleTemplate RuleTemplate
	if err := s.r.Request(ctx, http.MethodPatch, fmt.Sprintf("/cdn/resources/rule_templates/%d", id), req, &ruleTemplate); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &ruleTemplate, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if err := s.r.Request(ctx, http.MethodDelete, fmt.Sprintf("/cdn/resources/rule_templates/%d", id), nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
