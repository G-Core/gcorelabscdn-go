package clients

import (
	"context"
)

type ClientsMeService interface {
	Get(ctx context.Context) (*ClientsMe, error)
}

type ClientsMe struct {
	ID    int64  `json:"id"`
	Cname string `json:"cname"`
}
