package gcdn

import (
	"github.com/skyeng/gcorelabscdn-go/gcore"
	"github.com/skyeng/gcorelabscdn-go/origingroups"
	"github.com/skyeng/gcorelabscdn-go/resources"
	"github.com/skyeng/gcorelabscdn-go/rules"
	"github.com/skyeng/gcorelabscdn-go/sslcerts"
)

type ClientService interface {
	Resources() resources.ResourceService
	Rules() rules.RulesService
	OriginGroups() origingroups.OriginGroupService
	SSLCerts() sslcerts.SSLCertService
}

var _ ClientService = (*Service)(nil)

type Service struct {
	r                   gcore.Requester
	resourcesService    resources.ResourceService
	rulesService        rules.RulesService
	originGroupsService origingroups.OriginGroupService
	sslCertsService     sslcerts.SSLCertService
}

func NewService(r gcore.Requester) *Service {
	return &Service{
		r:                   r,
		resourcesService:    resources.NewService(r),
		rulesService:        rules.NewService(r),
		originGroupsService: origingroups.NewService(r),
		sslCertsService:     sslcerts.NewService(r),
	}
}

func (s *Service) Resources() resources.ResourceService {
	return s.resourcesService
}

func (s *Service) Rules() rules.RulesService {
	return s.rulesService
}

func (s *Service) OriginGroups() origingroups.OriginGroupService {
	return s.originGroupsService
}

func (s *Service) SSLCerts() sslcerts.SSLCertService {
	return s.sslCertsService
}
