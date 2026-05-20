package auth

import (
	"fmt"

	"github.com/hashicorp/go-retryablehttp"
)

// Func mutates a request before it is sent (e.g. Authorization header).
type Func func(*retryablehttp.Request) error

// Authenticator holds credentials for [Func] helpers.
type Authenticator struct {
	secret string
}

// New returns an authenticator with the given secret/password.
func New(secret string) *Authenticator {
	return &Authenticator{secret: secret}
}

// BearAuth returns Bearer token auth.
func (a *Authenticator) BearAuth() Func {
	return func(req *retryablehttp.Request) error {
		if a.secret == "" {
			return fmt.Errorf("httpclient/auth: bearer secret is empty")
		}
		req.Header.Set("Authorization", "Bearer "+a.secret)
		return nil
	}
}

// BasicAuth returns HTTP basic auth.
func (a *Authenticator) BasicAuth(username string) Func {
	return func(req *retryablehttp.Request) error {
		if username == "" || a.secret == "" {
			return fmt.Errorf("httpclient/auth: basic auth username or password is empty")
		}
		req.SetBasicAuth(username, a.secret)
		return nil
	}
}
