package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SolaTyolo/httpclient/response"
	"github.com/hashicorp/go-retryablehttp"
)

// Get sends a GET request. data is optional JSON body (unusual for GET; pass nil to omit).
func (h *HTTP) Get(ctx context.Context, endpoint string, data any) (*response.Response, error) {
	return h.request(ctx, http.MethodGet, endpoint, data)
}

// Post sends a POST request with an optional JSON body (nil for empty body).
func (h *HTTP) Post(ctx context.Context, endpoint string, data any) (*response.Response, error) {
	return h.request(ctx, http.MethodPost, endpoint, data)
}

// Put sends a PUT request with an optional JSON body.
func (h *HTTP) Put(ctx context.Context, endpoint string, data any) (*response.Response, error) {
	return h.request(ctx, http.MethodPut, endpoint, data)
}

// Delete sends a DELETE request with an optional JSON body.
func (h *HTTP) Delete(ctx context.Context, endpoint string, data any) (*response.Response, error) {
	return h.request(ctx, http.MethodDelete, endpoint, data)
}

func (h *HTTP) request(ctx context.Context, method, endpoint string, data any) (*response.Response, error) {
	var reader io.Reader
	if data != nil {
		payload, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("httpclient: marshal request: %w", err)
		}
		reader = bytes.NewReader(payload)
	}

	req, err := retryablehttp.NewRequestWithContext(ctx, method, endpoint, reader)
	if err != nil {
		return nil, fmt.Errorf("httpclient: create request: %w", err)
	}
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if h.authenticator != nil {
		if err := h.authenticator(req); err != nil {
			return nil, err
		}
	}

	resp, err := h.requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpclient: do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, response.StatusErrorFromReader(resp.StatusCode, resp.Body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("httpclient: read body: %w", err)
	}

	return &response.Response{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		RawBody:    body,
	}, nil
}
