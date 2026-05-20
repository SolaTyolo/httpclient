package response

import (
	"errors"
	"strings"
	"testing"
)

func TestStatusError(t *testing.T) {
	err := StatusErrorFromReader(404, strings.NewReader(`{"msg":"not found"}`))
	var se *StatusError
	if !errors.As(err, &se) {
		t.Fatalf("errors.As: %T", err)
	}
	if se.StatusCode != 404 || string(se.Body) != `{"msg":"not found"}` {
		t.Fatalf("got code=%d body=%q", se.StatusCode, se.Body)
	}
}
