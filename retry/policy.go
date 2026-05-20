package retry

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// DefaultRetryFn retries on HTTP 5xx responses.
func DefaultRetryFn(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	if resp != nil && resp.StatusCode >= 500 {
		return true, nil
	}
	return false, nil
}

// NoCheckRetryFn disables retries.
func NoCheckRetryFn(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return false, err
}

// Config holds retryablehttp client settings.
type Config struct {
	Timeout      time.Duration
	RetryMax     int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
	CheckRetry   retryablehttp.CheckRetry
}

// DefaultConfig returns package default retry settings.
func DefaultConfig() Config {
	return Config{
		Timeout:      DefaultTimeout,
		RetryMax:     DefaultRetryMax,
		RetryWaitMin: DefaultRetryWaitMin,
		RetryWaitMax: DefaultRetryWaitMax,
		CheckRetry:   DefaultRetryFn,
	}
}

// ApplyDefaults fills zero fields from [DefaultConfig].
func (c *Config) ApplyDefaults() {
	d := DefaultConfig()
	if c.Timeout <= 0 {
		c.Timeout = d.Timeout
	}
	if c.RetryMax <= 0 {
		c.RetryMax = d.RetryMax
	}
	if c.RetryWaitMin <= 0 {
		c.RetryWaitMin = d.RetryWaitMin
	}
	if c.RetryWaitMax <= 0 {
		c.RetryWaitMax = d.RetryWaitMax
	}
	if c.CheckRetry == nil {
		c.CheckRetry = d.CheckRetry
	}
}

// NewClient builds a *retryablehttp.Client from cfg (after [Config.ApplyDefaults]).
func NewClient(cfg Config) *retryablehttp.Client {
	c := retryablehttp.NewClient()
	c.RetryMax = cfg.RetryMax
	c.RetryWaitMin = cfg.RetryWaitMin
	c.RetryWaitMax = cfg.RetryWaitMax
	c.CheckRetry = cfg.CheckRetry
	if cfg.Timeout > 0 {
		c.HTTPClient.Timeout = cfg.Timeout
	}
	return c
}
