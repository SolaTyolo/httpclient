package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	DefaultRetryWaitMin = 2 * time.Second
	DefaultRetryWaitMax = 10 * time.Second
	DefaultRetryMax     = 3
	DefaultTimeout      = 5 * time.Second
)

type HTTP struct {
	key           string
	requester     Requester
	authenticator AuthFunc
}

type Requester interface {
	Do(*retryablehttp.Request) (*http.Response, error)
}

func Default() *HTTP {
	return New(WithHTTPRequester(newDefaultRequester()))
}

func newDefaultRequester() *retryablehttp.Client {
	client := retryablehttp.NewClient()

	client.RetryMax = DefaultRetryMax
	client.RetryWaitMin = DefaultRetryWaitMin
	client.RetryWaitMax = DefaultRetryWaitMax
	client.CheckRetry = DefaultRetryFn

	client.HTTPClient.Timeout = DefaultTimeout
	return client
}

func New(opts ...ClientOption) *HTTP {
	h := &HTTP{}
	for _, opt := range opts {
		opt(h)
	}

	if h.requester == nil {
		h.requester = newDefaultRequester()
	}

	return h
}

func (h *HTTP) Get(ctx context.Context, endpoint string, data any) (*HttpResponse, error) {
	return h.request(ctx, http.MethodGet, endpoint, data)
}
func (h *HTTP) Post(ctx context.Context, endpoint string, data any) (*HttpResponse, error) {
	return h.request(ctx, http.MethodPost, endpoint, data)
}
func (h *HTTP) Put(ctx context.Context, endpoint string, data any) (*HttpResponse, error) {
	return h.request(ctx, http.MethodPut, endpoint, data)
}
func (h *HTTP) Delete(ctx context.Context, endpoint string, data any) (*HttpResponse, error) {
	return h.request(ctx, http.MethodDelete, endpoint, data)
}

func (h *HTTP) request(ctx context.Context, method string, endpoint string, data any) (*HttpResponse, error) {
	var reader io.Reader
	if data != nil {
		payload, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("cannot marshal request: %w", err)
		}
		reader = bytes.NewReader(payload)
	}

	req, err := retryablehttp.NewRequestWithContext(ctx, method, endpoint, reader)
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}
	// TODO -- Extensions SetHeader
	req.Header.Set("Content-Type", "application/json")

	if h.authenticator != nil {
		if err := h.authenticator(req); err != nil {
			return nil, err
		}
	}

	resp, err := h.requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		err := h.makeClientError(resp.StatusCode, resp.Body)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response: %w", err)
	}

	return &HttpResponse{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		RawBody:    body,
	}, nil
}

func (h *HTTP) makeClientError(statusCode int, body io.Reader) error {
	if body == nil {
		return errors.New("invalid body")
	}
	errBody, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return fmt.Errorf("unexpected error (status code %d), err:%s", statusCode, string(errBody))
}

type HttpResponse struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	RawBody    []byte      `json:"-"`
}

func (resp HttpResponse) String() string {
	contentType := resp.Header.Get("Content-Type")
	body := fmt.Sprintf("<binary> len %d", len(resp.RawBody))
	if strings.Contains(contentType, "json") || strings.Contains(contentType, "text") {
		body = string(resp.RawBody)
	}
	return fmt.Sprintf("StatusCode: %d, Header:%v, Content-Type: %s, Body: %v", resp.StatusCode,
		resp.Header, resp.Header.Get("Content-Type"), body)
}
