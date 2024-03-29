package originshielding

import (
	"context"
	"fmt"
	"github.com/G-Core/gcorelabscdn-go/gcore"
	"net/http"
)

var _ OriginShieldingService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Update(ctx context.Context, id int64, req *UpdateRequest) (*OriginShieldingData, error) {
	var originShielding OriginShieldingData
	if err := s.r.Request(ctx, http.MethodPut, fmt.Sprintf("/cdn/resources/%d/shielding_v2", id), req, &originShielding); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return &originShielding, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*OriginShieldingData, error) {
	var originShielding OriginShieldingData
	if err := s.r.Request(ctx, http.MethodGet, fmt.Sprintf("/cdn/resources/%d/shielding_v2", id), nil, &originShielding); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &originShielding, nil

}

func (s *Service) GetLocations(ctx context.Context) (*[]OriginShieldingLocations, error) {
	var originShieldingLocations []OriginShieldingLocations
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/shieldingpop_v2", nil, &originShieldingLocations); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &originShieldingLocations, nil
}
