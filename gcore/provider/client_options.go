package provider

import (
	"time"

	"github.com/skyeng/gcorelabscdn-go/gcore"
)

type ClientOption func(*Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpc.Timeout = timeout
	}
}

func WithSigner(signer gcore.RequestSigner) ClientOption {
	return func(c *Client) {
		c.signer = signer
	}
}

func WithSignerFunc(f gcore.RequestSignerFunc) ClientOption {
	return func(c *Client) {
		c.signer = f
	}
}

func WithUA(ua string) ClientOption {
	return func(c *Client) {
		c.ua = ua
	}
}
