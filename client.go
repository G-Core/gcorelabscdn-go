package gcdn

import (
	"github.com/G-Core/gcorelabscdn-go/cacerts"
	"github.com/G-Core/gcorelabscdn-go/gcore"
	"github.com/G-Core/gcorelabscdn-go/origingroups"
	"github.com/G-Core/gcorelabscdn-go/originshielding"
	"github.com/G-Core/gcorelabscdn-go/presets"
	"github.com/G-Core/gcorelabscdn-go/resources"
	"github.com/G-Core/gcorelabscdn-go/rules"
	"github.com/G-Core/gcorelabscdn-go/sslcerts"
)

type ClientService interface {
	Resources() resources.ResourceService
	Rules() rules.RulesService
	OriginGroups() origingroups.OriginGroupService
	OriginShielding() originshielding.OriginShieldingService
	SSLCerts() sslcerts.SSLCertService
	Presets() presets.PresetsService
	CACerts() cacerts.CACertService
}

var _ ClientService = (*Service)(nil)

type Service struct {
	r                      gcore.Requester
	resourcesService       resources.ResourceService
	rulesService           rules.RulesService
	originGroupsService    origingroups.OriginGroupService
	originShieldingService originshielding.OriginShieldingService
	sslCertsService        sslcerts.SSLCertService
	presetsService         presets.PresetsService
	caCertsService         cacerts.CACertService
}

func NewService(r gcore.Requester) *Service {
	return &Service{
		r:                      r,
		resourcesService:       resources.NewService(r),
		rulesService:           rules.NewService(r),
		originGroupsService:    origingroups.NewService(r),
		originShieldingService: originshielding.NewService(r),
		sslCertsService:        sslcerts.NewService(r),
		presetsService:         presets.NewService(r),
		caCertsService:         cacerts.NewService(r),
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

func (s *Service) OriginShielding() originshielding.OriginShieldingService {
	return s.originShieldingService
}

func (s *Service) SSLCerts() sslcerts.SSLCertService {
	return s.sslCertsService
}

func (s *Service) Presets() presets.PresetsService {
	return s.presetsService
}

func (s *Service) CACerts() cacerts.CACertService {
	return s.caCertsService
}
