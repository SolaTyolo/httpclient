package response

import "testing"

func TestResponse_DecodeJSON(t *testing.T) {
	var out struct {
		OK bool `json:"ok"`
	}
	resp := &Response{RawBody: []byte(`{"ok":true}`)}
	if err := resp.DecodeJSON(&out); err != nil {
		t.Fatal(err)
	}
	if !out.OK {
		t.Fatal("expected ok")
	}
}
