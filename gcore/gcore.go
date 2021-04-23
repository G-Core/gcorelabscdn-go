package gcore

import (
	"context"
	"net/http"
)

type Requester interface {
	Request(ctx context.Context, method, path string, payload interface{}, result interface{}) error
}

type RequestSigner interface {
	Sign(req *http.Request) error
}

type RequestSignerFunc func(req *http.Request) error

func (rsf RequestSignerFunc) Sign(req *http.Request) error {
	return rsf(req)
}
