package test_test

import (
	"github.com/dooman87/kolibri/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleRunRequests() {
	//This is service that we want to test. It should be a type of http.HandlerFunc.
	test.Service = func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("OK"))
	}

	//Creating set of tests that we want to run.
	//Each test is a struct that contains endpoint, expected response status, description, optional handler.
	testCases := []test.TestCase{
		{"http://localhost:8080", http.StatusOK, "Should return 200", nil},
		{"http://localhost:8080", http.StatusOK, "Should return OK in body", func(w *httptest.ResponseRecorder, t *testing.T) {
			if w.Body.String() != "OK" {
				t.Errorf("Expected %s but got %s", "OK", w.Body.String())
			}
		}},
	}

	//Running all test cases.
	test.RunRequests(testCases, &testing.T{})
}
