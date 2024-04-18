package httpclient

import "testing"

func TestMakeEndpoint(t *testing.T) {
	host := "http://example.com"
	format := "/path/%s/%d"
	expectedPath := "/path/value1/123"
	expectedURL := host + expectedPath

	endpoint := MakeEndpoint(host, format, "value1", 123)

	if endpoint.url.String() != expectedURL {
		t.Errorf("Incorrect URL, got: %s, want: %s.", endpoint.url.String(), expectedURL)
	}

	if len(endpoint.query) != 0 {
		t.Errorf("Query values should be empty, got: %v.", endpoint.query)
	}
}

func TestMakeEndpointWithoutFormat(t *testing.T) {
	host := "http://example.com"
	format := ""
	expectedURL := host

	endpoint := MakeEndpoint(host, format)
	if endpoint.String() != expectedURL {
		t.Errorf("Incorrect URL, got: %s, want: %s.", endpoint.url.String(), expectedURL)
	}
}
