package endpoint

import (
	"fmt"
	"net/url"
)

// Make builds a URL from host + path format (fmt.Sprintf) and optional query params.
func Make(host, format string, args ...any) (Endpoint, error) {
	path := fmt.Sprintf(format, args...)
	u, err := url.Parse(host + path)
	if err != nil {
		return Endpoint{}, fmt.Errorf("httpclient/endpoint: parse endpoint: %w", err)
	}
	return Endpoint{url: u, query: make(url.Values)}, nil
}

// Endpoint is a URL builder with query parameters.
type Endpoint struct {
	url   *url.URL
	query url.Values
}

func (e Endpoint) String() string {
	u := *e.url
	u.RawQuery = e.query.Encode()
	return u.String()
}

// AddQueryParam adds a query parameter when valuer is valid.
func (e *Endpoint) AddQueryParam(v Valuer) {
	if v == nil || !v.Valid() {
		return
	}
	k, val := v.Values()
	e.query.Add(k, val)
}

// Valuer is a single query key/value.
type Valuer interface {
	Values() (key, value string)
	Valid() bool
}

// RequestOption marks types usable as query parameters.
type RequestOption interface {
	Valuer
}

type baseRequestOption struct {
	key   string
	value string
}

// MakeRequestOption builds a query parameter option.
func MakeRequestOption(key string, value any) RequestOption {
	return baseRequestOption{key: key, value: fmt.Sprintf("%v", value)}
}

func (o baseRequestOption) Values() (string, string) { return o.key, o.value }
func (o baseRequestOption) Valid() bool             { return o.value != "" }
