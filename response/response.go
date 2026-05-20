package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Response is the result of a successful JSON API call (2xx).
type Response struct {
	StatusCode int
	Header     http.Header
	RawBody    []byte
}

// DecodeJSON unmarshals RawBody into v.
func (resp *Response) DecodeJSON(v any) error {
	if len(resp.RawBody) == 0 {
		return fmt.Errorf("httpclient/response: empty response body")
	}
	if err := json.Unmarshal(resp.RawBody, v); err != nil {
		return fmt.Errorf("httpclient/response: decode json: %w", err)
	}
	return nil
}

func (resp Response) String() string {
	contentType := resp.Header.Get("Content-Type")
	body := fmt.Sprintf("<binary> len %d", len(resp.RawBody))
	if strings.Contains(contentType, "json") || strings.Contains(contentType, "text") {
		body = string(resp.RawBody)
	}
	return fmt.Sprintf("StatusCode: %d, Header: %v, Content-Type: %s, Body: %v",
		resp.StatusCode, resp.Header, contentType, body)
}
