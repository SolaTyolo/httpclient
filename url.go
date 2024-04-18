package httpclient

import (
	"fmt"
	"net/url"
)

func MakeEndpoint(host string, format string, a ...any) endpoint {
	path := fmt.Sprintf(format, a...)
	u, _ := url.Parse(host + path)

	return endpoint{
		url:   u,
		query: make(url.Values),
	}
}

type endpoint struct {
	url   *url.URL
	query url.Values
}

func (e endpoint) String() string {
	e.url.RawQuery = e.query.Encode()
	return e.url.String()
}

func (e *endpoint) AddQueryParam(v valuer) {
	if !v.valid() {
		return
	}
	e.query.Add(v.values())
}

type valuer interface {
	values() (string, string)
	valid() bool
}

// requestOption is an interface representing API request optional filters and
type requestOption interface {
	valuer
}

type baseRequestOption struct {
	key   string
	value string
}

func MakeRequestOption(key string, value any) requestOption {
	return baseRequestOption{
		key:   key,
		value: fmt.Sprintf("%v", value),
	}
}

func (o baseRequestOption) values() (key, value string) {
	return o.key, o.value
}

func (o baseRequestOption) valid() bool {
	return o.value != ""
}
