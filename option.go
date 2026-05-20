package httpclient

import (
	"time"

	"github.com/SolaTyolo/httpclient/auth"
	"github.com/SolaTyolo/httpclient/retry"
	"github.com/SolaTyolo/httpclient/transport"
	"github.com/hashicorp/go-retryablehttp"
)

// ClientOption configures [HTTP] / [New].
type ClientOption func(*clientConfig)

type clientConfig struct {
	retry         retry.Config
	requester     transport.Requester
	authenticator auth.Func
}

// WithTimeout sets per-attempt timeout on the underlying retryablehttp client.
func WithTimeout(d time.Duration) ClientOption {
	return func(cfg *clientConfig) {
		cfg.retry.Timeout = d
	}
}

// WithRetryMax sets maximum retry attempts.
func WithRetryMax(n int) ClientOption {
	return func(cfg *clientConfig) {
		cfg.retry.RetryMax = n
	}
}

// WithRetryWait sets backoff bounds between retries.
func WithRetryWait(min, max time.Duration) ClientOption {
	return func(cfg *clientConfig) {
		cfg.retry.RetryWaitMin = min
		cfg.retry.RetryWaitMax = max
	}
}

// WithCheckRetry sets the retry policy (e.g. retry.DefaultRetryFn, retry.NoCheckRetryFn).
func WithCheckRetry(fn retryablehttp.CheckRetry) ClientOption {
	return func(cfg *clientConfig) {
		cfg.retry.CheckRetry = fn
	}
}

// WithAuthFn runs before each JSON API request.
func WithAuthFn(fn auth.Func) ClientOption {
	return func(cfg *clientConfig) {
		cfg.authenticator = fn
	}
}

// WithRequester sets the transport. Use [WithRetryableClient] for the common case.
func WithRequester(r transport.Requester) ClientOption {
	return func(cfg *clientConfig) {
		cfg.requester = r
	}
}

// WithRetryableClient sets a *retryablehttp.Client as the transport.
func WithRetryableClient(c *retryablehttp.Client) ClientOption {
	return WithRequester(c)
}
