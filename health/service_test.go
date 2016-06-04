package health_test

import (
	"github.com/dooman87/kolibri/health"
	"github.com/dooman87/kolibri/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	test.Service = health.Health

	testCases := []test.TestCase{
		{"http://localhost/health", http.StatusOK, "Should return OK in body", func(w *httptest.ResponseRecorder, t *testing.T) {
			expectedResponse := "OK"
			if w.Body.String() != expectedResponse {
				t.Errorf("Expected %s but got %s", expectedResponse, w.Body.String())
			}
		}},
	}
	test.RunRequests(testCases, t)
}
