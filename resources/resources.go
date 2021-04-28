package resources

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type ResourceService interface {
	Create(ctx context.Context, req *CreateRequest) (*Resource, error)
	Get(ctx context.Context, id int64) (*Resource, error)
	Update(ctx context.Context, id int64, req *UpdateRequest) (*Resource, error)
}

var _ ResourceService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodPost, "/resources", req, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/resources/%d", id), nil, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}

func (s *Service) Update(ctx context.Context, id int64, req *UpdateRequest) (*Resource, error) {
	var resource Resource
	if err := s.r.Request(ctx, http.MethodPut, fmt.Sprintf("/resources/%d", id), req, &resource); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resource, nil
}

type Protocol string

const (
	HTTPProtocol   Protocol = "HTTP"
	HTTPSProtocol  Protocol = "HTTPS"
	MatchProtocaol Protocol = "MATCH"
)

type ResourceStatus string

const (
	ActiveResourceStatus    ResourceStatus = "Active"
	SuspendedResourceStatus ResourceStatus = "Suspended"
	ProcessedResourceStatus ResourceStatus = "Processed"
)

type CreateRequest struct {
	Cname              string   `json:"cname,omitempty"`
	OriginGroup        int      `json:"originGroup,omitempty"`
	Origin             string   `json:"origin,omitempty"`
	SecondaryHostnames []string `json:"secondaryHostnames,omitempty"`
}

type UpdateRequest struct {
	Active             bool     `json:"active"`
	OriginGroup        int      `json:"originGroup"`
	SecondaryHostnames []string `json:"secondaryHostnames,omitempty"`
	SSlEnabled         bool     `json:"sslEnabled"`
	SSLData            int      `json:"sslData,omitempty"`
	SSLAutomated       bool     `json:"ssl_automated"`
	OriginProtocol     Protocol `json:"originProtocol,omitempty"`
}

type Resource struct {
	ID                 int64          `json:"id"`
	Name               string         `json:"name"`
	CreatedAt          time.Time      `json:"created"`
	UpdatedAt          time.Time      `json:"updated"`
	Status             ResourceStatus `json:"stratus"`
	Active             bool           `json:"active"`
	Client             int64          `json:"client"`
	OriginGroup        int64          `json:"originGroup"`
	Cname              string         `json:"cname"`
	SecondaryHostnames []string       `json:"secondaryHostnames"`
	Shielded           bool           `json:"shielded"`
	SSlEnabled         bool           `json:"sslEnabled"`
	SSLData            int            `json:"sslData"`
	SSLAutomated       bool           `json:"ssl_automated"`
	OriginProtocol     Protocol       `json:"originProtocol"`
}
