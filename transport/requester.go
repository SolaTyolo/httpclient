package transport

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

// Requester performs HTTP requests. Production code should use *retryablehttp.Client.
type Requester interface {
	Do(*retryablehttp.Request) (*http.Response, error)
}
