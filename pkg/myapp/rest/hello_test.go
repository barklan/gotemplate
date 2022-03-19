package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	if testing.Short() {
		return
	}
	t.Parallel()
	c, err := MockCtrl()
	if err != nil {
		t.Errorf("failed to construct mock controller: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/?", nil)
	w := httptest.NewRecorder()
	c.helloHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("status code is %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	resp := make(map[string]string)
	if err = json.Unmarshal(body, &resp); err != nil {
		t.Errorf("failed to unmarshal resp: %v", err)
	}
	if _, ok := resp["hello"]; !ok {
		t.Errorf("expected to have hello key")
	}
}
