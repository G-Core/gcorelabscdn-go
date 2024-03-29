package rules

import (
	"context"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type RulesService interface {
	Create(ctx context.Context, resourceID int64, req *CreateRequest) (*Rule, error)
	Get(ctx context.Context, resourceID, ruleID int64) (*Rule, error)
	Update(ctx context.Context, resourceID, ruleID int64, req *UpdateRequest) (*Rule, error)
	Delete(ctx context.Context, resourceID, ruleID int64) error
}

type CreateRequest struct {
	Name                   string         `json:"name,omitempty"`
	Active                 bool           `json:"active"`
	Rule                   string         `json:"rule,omitempty"`
	RuleType               int            `json:"ruleType"`
	OriginGroup            *int           `json:"originGroup"`
	Weight                 int            `json:"weight,omitempty"`
	OverrideOriginProtocol *string        `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}

type UpdateRequest struct {
	Name                   string         `json:"name,omitempty"`
	Active                 bool           `json:"active"`
	Rule                   string         `json:"rule,omitempty"`
	RuleType               int            `json:"ruleType"`
	Weight                 int            `json:"weight,omitempty"`
	OriginGroup            *int           `json:"originGroup"`
	OverrideOriginProtocol *string        `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}

type ActivateRequest struct {
	Active bool `json:"active"`
}

type Rule struct {
	ID                     int64          `json:"id"`
	Name                   string         `json:"name"`
	Active                 bool           `json:"active"`
	Deleted                bool           `json:"deleted"`
	OriginGroup            *int           `json:"originGroup"`
	OriginProtocol         string         `json:"originProtocol"`
	OverrideOriginProtocol *string        `json:"overrideOriginProtocol"`
	Pattern                string         `json:"rule"`
	Type                   int            `json:"ruleType"`
	Weight                 int            `json:"weight"`
	Options                *gcore.Options `json:"options,omitempty"`
}
