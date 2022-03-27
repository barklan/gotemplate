//go:build e2e
// +build e2e

package myapp

import (
	"net/http"
	"testing"
)

func Test_myapp(t *testing.T) {
	_, err := http.Get("https://google.com")
	if err != nil {
		t.Fatal(err)
	}
}
