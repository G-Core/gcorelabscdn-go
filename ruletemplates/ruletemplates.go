package ruletemplates

import (
	"context"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type RuleTemplateService interface {
	Create(ctx context.Context, req *CreateRequest) (*RuleTemplate, error)
	Get(ctx context.Context, id int64) (*RuleTemplate, error)
	Update(ctx context.Context, id int64, req *UpdateRequest) (*RuleTemplate, error)
	Delete(ctx context.Context, id int64) error
}

type Protocol string

const (
	HTTPProtocol  Protocol = "HTTP"
	HTTPSProtocol Protocol = "HTTPS"
	MatchProtocol Protocol = "MATCH"
)

type RuleType int

const (
	RuleType0 RuleType = 0
	RuleType1 RuleType = 1
)

type RuleTemplate struct {
	ID                     int64          `json:"id"`
	Name                   string         `json:"name"`
	Client                 int64          `json:"client"`
	Deleted                bool           `json:"deleted"`
	Rule                   string         `json:"rule"`
	RuleType               RuleType       `json:"ruleType"`
	Weight                 int            `json:"weight"`
	Template               bool           `json:"template"`
	Default                bool           `json:"default"`
	OverrideOriginProtocol *Protocol      `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}

type CreateRequest struct {
	Name                   string         `json:"name,omitempty"`
	Rule                   string         `json:"rule"`
	RuleType               RuleType       `json:"ruleType"`
	Weight                 int            `json:"weight,omitempty"`
	OverrideOriginProtocol *Protocol      `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}

type UpdateRequest struct {
	Name                   string         `json:"name,omitempty"`
	Rule                   string         `json:"rule,omitempty"`
	RuleType               RuleType       `json:"ruleType,omitempty"`
	Weight                 int            `json:"weight,omitempty"`
	OverrideOriginProtocol *Protocol      `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}
