package retry

import (
	"context"
	"testing"
)

func TestDefaultRetryFn(t *testing.T) {
	retry, err := DefaultRetryFn(context.Background(), nil, nil)
	if retry || err != nil {
		t.Fatalf("retry=%v err=%v", retry, err)
	}
}
