package resources

import (
	"context"
	"time"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type ResourceService interface {
	Create(ctx context.Context, req *CreateRequest) (*Resource, error)
	Get(ctx context.Context, id int64) (*Resource, error)
	Update(ctx context.Context, id int64, req *UpdateRequest) (*Resource, error)
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
	Description        string   `json:"description,omitempty"`
}

type UpdateRequest struct {
	Active             bool           `json:"active"`
	OriginGroup        int            `json:"originGroup"`
	SecondaryHostnames []string       `json:"secondaryHostnames,omitempty"`
	SSlEnabled         bool           `json:"sslEnabled"`
	SSLData            int            `json:"sslData,omitempty"`
	SSLAutomated       bool           `json:"ssl_automated"`
	OriginProtocol     Protocol       `json:"originProtocol,omitempty"`
	Options            *gcore.Options `json:"options,omitempty"`
	Description        string         `json:"description,omitempty"`
}

type Resource struct {
	ID                 int64          `json:"id"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	CreatedAt          time.Time      `json:"created"`
	UpdatedAt          time.Time      `json:"updated"`
	Status             ResourceStatus `json:"status"`
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
	Options            *gcore.Options `json:"options"`
}
