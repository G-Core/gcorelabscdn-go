package clients

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Get(ctx context.Context) (*ClientsMe, error) {
	var client ClientsMe
	if err := s.r.Request(ctx, http.MethodGet, "/cdn/clients/me", nil, &client); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &client, nil
}
