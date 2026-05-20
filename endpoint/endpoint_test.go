package endpoint

import "testing"

func TestMake(t *testing.T) {
	ep, err := Make("http://example.com", "/path/%s/%d", "value1", 123)
	if err != nil {
		t.Fatal(err)
	}
	want := "http://example.com/path/value1/123"
	if got := ep.String(); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMakeWithoutFormat(t *testing.T) {
	ep, err := Make("http://example.com", "")
	if err != nil {
		t.Fatal(err)
	}
	if got := ep.String(); got != "http://example.com" {
		t.Errorf("got %q", got)
	}
}
