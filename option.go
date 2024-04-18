package httpclient

import (
	"context"
	"net/http"
)

type ClientOption func(*HTTP)

func WithKey(key string) ClientOption {
	return func(h *HTTP) {
		h.key = key
	}
}

// WithHTTPRequester sets the HTTP requester for a given client, used mostly for testing.
func WithHTTPRequester(requester Requester) ClientOption {
	return func(c *HTTP) {
		c.requester = requester
	}
}

func WithAuthFn(fn AuthFunc) ClientOption {
	return func(c *HTTP) {
		c.authenticator = fn
	}
}

func NoCheckRetryFn(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return false, err
}

func DefaultRetryFn(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 500 {
		return true, nil
	}
	return false, err
}
