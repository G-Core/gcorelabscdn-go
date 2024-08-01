package invalidations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ InvalidationsService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) PurgeByPattern(ctx context.Context, resourceID int, req *PurgeByPatternRequest) error {
	path := fmt.Sprintf("/cdn/resources/%d/purge", resourceID)
	if err := s.r.Request(ctx, http.MethodPost, path, req, nil); err != nil {
		return fmt.Errorf("received error from gcore: %w", err)
	}

	return nil
}
