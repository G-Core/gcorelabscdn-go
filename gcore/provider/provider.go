package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

type Client struct {
	httpc   *http.Client
	signer  gcore.RequestSigner
	ua      string
	baseURL string
}

func NewClient(baseURL string, opts ...ClientOption) *Client {
	httpc := &http.Client{Timeout: time.Minute}
	c := &Client{httpc: httpc, baseURL: baseURL}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Request(ctx context.Context, method, path string, payload interface{}, result interface{}) error {
	var body io.Reader
	if payload != nil {
		payloadBuf := new(bytes.Buffer)
		if err := json.NewEncoder(payloadBuf).Encode(payload); err != nil {
			return fmt.Errorf("encode request payload: %w", err)
		}

		body = payloadBuf
	}

	// TODO: figure out how to drop trailing slash
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return fmt.Errorf("create HTTP request: %w", err)
	}

	resp, err := c.do(req)
	if err != nil {
		return fmt.Errorf("execute HTTP request: %w", err)
	}

	rawBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errResp gcore.ErrorResponse
		errResp.StatusCode = resp.StatusCode

		if err := json.Unmarshal(rawBody, &errResp); err != nil {
			return fmt.Errorf("unexpected response, status: %d, body: %s", resp.StatusCode, string(rawBody))
		}

		return &errResp
	}

	if result != nil {
		if err := json.Unmarshal(rawBody, result); err != nil {
			return fmt.Errorf("unable to decode response, status: %d, body: %s", resp.StatusCode, string(rawBody))
		}
	}

	return nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	if c.ua != "" {
		req.Header.Set("User-Agent", c.ua)
	}
	if c.signer != nil {
		if err := c.signer.Sign(req); err != nil {
			return nil, err
		}
	}

	return c.httpc.Do(req)
}
