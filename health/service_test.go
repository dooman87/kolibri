package health_test

import (
	"github.com/dooman87/kolibri/health"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	health.Health(w, req)

	expected := "OK"
	if w.Body.String() != expected {
		t.Fatalf("Expected %s but got %s", expected, w.Body.String())
	}
}
