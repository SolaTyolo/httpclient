package httpclient

import (
	"fmt"

	"github.com/hashicorp/go-retryablehttp"
)

type AuthFunc func(*retryablehttp.Request) error

type Authenticator struct {
	secret string
}

func NewAuthenticator(secret string) *Authenticator {
	return &Authenticator{
		secret: secret,
	}
}

func (a Authenticator) BearAuth() AuthFunc {
	return func(req *retryablehttp.Request) error {
		if a.secret == "" {
			return fmt.Errorf("[HTTP] secret is empty")
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.secret))
		return nil
	}
}

func (a Authenticator) BasicAuth(key string) AuthFunc {
	return func(req *retryablehttp.Request) error {
		if key == "" || a.secret == "" {
			return fmt.Errorf("[HTTP] key or secret is empty")
		}
		req.SetBasicAuth(key, a.secret)
		return nil
	}
}
