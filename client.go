package gcdn

import (
	"github.com/G-Core/gcorelabscdn-go/gcore"
	"github.com/G-Core/gcorelabscdn-go/resources"
	"github.com/G-Core/gcorelabscdn-go/rules"
)

type ClientService interface {
	Resources() resources.ResourceService
	Rules() rules.RulesService
}

var _ ClientService = (*Service)(nil)

type Service struct {
	r                gcore.Requester
	resourcesService resources.ResourceService
	rulesService     rules.RulesService
}

func NewService(baseURL string, r gcore.Requester) *Service {
	return &Service{
		r:                r,
		resourcesService: resources.NewService(r),
		rulesService:     rules.NewService(r),
	}
}

func (s *Service) Resources() resources.ResourceService {
	return s.resourcesService
}

func (s *Service) Rules() rules.RulesService {
	return s.rulesService
}
