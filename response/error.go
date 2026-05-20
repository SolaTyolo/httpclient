package response

import (
	"fmt"
	"io"
)

// StatusError is returned for non-2xx responses. Use errors.As to inspect StatusCode and Body.
type StatusError struct {
	StatusCode int
	Body       []byte
}

func (e *StatusError) Error() string {
	if len(e.Body) == 0 {
		return fmt.Sprintf("httpclient/response: unexpected status %d", e.StatusCode)
	}
	return fmt.Sprintf("httpclient/response: unexpected status %d: %s", e.StatusCode, string(e.Body))
}

// StatusErrorFromReader builds a StatusError from an HTTP error response body.
func StatusErrorFromReader(code int, body io.Reader) error {
	if body == nil {
		return &StatusError{StatusCode: code}
	}
	b, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("httpclient/response: read error body: %w", err)
	}
	return &StatusError{StatusCode: code, Body: b}
}
