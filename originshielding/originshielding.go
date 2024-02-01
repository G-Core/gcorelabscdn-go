package originshielding

import (
	"context"
)

type OriginShieldingService interface {
	Update(ctx context.Context, id int64, req *UpdateRequest) (*OriginShieldingData, error)
	Get(ctx context.Context, id int64) (*OriginShieldingData, error)
	GetLocations(ctx context.Context) (*OriginShieldingLocations, error)
}

type OriginShieldingData struct {
	ShieldingPop int `json:"shielding_pop"`
}

type UpdateRequest struct {
	ShieldingPop int `json:"shielding_pop"`
}

type OriginShieldingLocations struct {
	ID         int    `json:"id"`
	Datacenter string `json:"datacenter"`
	Country    string `json:"country"`
	City       string `json:"city"`
}
