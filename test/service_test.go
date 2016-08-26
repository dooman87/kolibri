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
	test.T = &testing.T{}

	//Creating set of tests that we want to run.
	//Each test is a struct that contains endpoint, expected response status, description, optional handler.
	testCases := []test.TestCase{
		{
			Url:         "http://localhost:8080",
			Description: "Should return 200",
		},
		{
			Url:         "http://localhost:8080",
			Description: "Should return OK in body",
			Handler: func(w *httptest.ResponseRecorder, t *testing.T) {
				if w.Body.String() != "OK" {
					t.Errorf("Expected %s but got %s", "OK", w.Body.String())
				}
			},
		},
		{
			Request:     test.NewRequest("GET", "http://localhost:8080", nil),
			Description: "Should return 200",
		},
	}

	//Running all test cases.
	test.RunRequests(testCases)
}
