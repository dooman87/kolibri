package test_test

import (
	"github.com/dooman87/kolibri/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleRunRequests() {
	test.Service = func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("OK"))
	}
	testCases := []test.TestCase{
		{"http://localhost:8080", http.StatusOK, "Should return 200", nil},
		{"http://localhost:8080", http.StatusOK, "Should return OK in body", func(w *httptest.ResponseRecorder, t *testing.T) {
			if w.Body.String() != "OK" {
				t.Errorf("Expected %s but got %s", "OK", w.Body.String())
			}
		}},
	}

	test.RunRequests(testCases, &testing.T{})
}
