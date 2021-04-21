package gcdn

import (
	"github.com/G-Core/gcorelabscdn-go/gcore"
	"github.com/G-Core/gcorelabscdn-go/resources"
)

type ClientService interface {
	Resources(opts ...resources.ServiceOption) resources.ResourceService
}

type Service struct {
	r       gcore.Requester
	baseURL string
}

func NewService(baseURL string, r gcore.Requester) *Service {
	return &Service{baseURL: baseURL, r: r}
}

func (s *Service) Path() string {
	return s.baseURL
}

func (s *Service) Resources(opts ...resources.ServiceOption) *resources.Service {
	return resources.NewService(s, s.r, opts...)
}
