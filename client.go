package httpclient

import (
	"github.com/SolaTyolo/httpclient/auth"
	"github.com/SolaTyolo/httpclient/retry"
	"github.com/SolaTyolo/httpclient/transport"
)

// HTTP is a retryable JSON HTTP client built on go-retryablehttp.
type HTTP struct {
	requester     transport.Requester
	authenticator auth.Func
}

// Default returns a client with package default retry and timeout settings.
func Default() *HTTP {
	return New()
}

// New builds an HTTP client.
func New(opts ...ClientOption) *HTTP {
	cfg := &clientConfig{retry: retry.DefaultConfig()}
	for _, opt := range opts {
		opt(cfg)
	}
	cfg.retry.ApplyDefaults()

	h := &HTTP{authenticator: cfg.authenticator}
	if cfg.requester != nil {
		h.requester = cfg.requester
	} else {
		h.requester = retry.NewClient(cfg.retry)
	}
	return h
}

// Requester returns the transport used by JSON API methods.
func (h *HTTP) Requester() transport.Requester {
	return h.requester
}
