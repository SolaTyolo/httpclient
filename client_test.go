package httpclient

import (
	"context"
	"testing"
	"time"

	"github.com/SolaTyolo/httpclient/retry"
	"github.com/hashicorp/go-retryablehttp"
)

func TestWithTimeout_appliedToRetryClient(t *testing.T) {
	h := New(WithTimeout(30 * time.Second))
	rc, ok := h.Requester().(*retryablehttp.Client)
	if !ok {
		t.Fatalf("requester type %T, want *retryablehttp.Client", h.Requester())
	}
	if rc.HTTPClient.Timeout != 30*time.Second {
		t.Fatalf("timeout: got %v want 30s", rc.HTTPClient.Timeout)
	}
}

func TestDefaultRetryFn(t *testing.T) {
	ok, err := retry.DefaultRetryFn(context.Background(), nil, nil)
	if ok || err != nil {
		t.Fatalf("retry=%v err=%v", ok, err)
	}
}
