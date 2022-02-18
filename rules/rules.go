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
	Rule                   string         `json:"rule,omitempty"`
	RuleType               int            `json:"ruleType"`
	OverrideOriginProtocol *string        `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}

type UpdateRequest struct {
	Name                   string         `json:"name,omitempty"`
	Rule                   string         `json:"rule,omitempty"`
	RuleType               int            `json:"ruleType"`
	OverrideOriginProtocol *string        `json:"overrideOriginProtocol"`
	Options                *gcore.Options `json:"options,omitempty"`
}

type Rule struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Deleted        bool           `json:"deleted"`
	OriginGroup    int            `json:"originGroup"`
	OriginProtocol string         `json:"originProtocol"`
	Pattern        string         `json:"rule"`
	Type           int            `json:"ruleType"`
	Options        *gcore.Options `json:"options,omitempty"`
}
