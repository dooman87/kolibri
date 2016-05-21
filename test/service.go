//Provides helper to test http handlers.
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//Describes test case. Each test case
//will be executed as http request agains
//Service handler.
type TestCase struct {
	Url          string
	ExpectedCode int
	Description  string
}

var (
	//Handler that will be executed for test cases.
	Service http.HandlerFunc
)

//Runs all test cases and failing test if
//response code is not equals to expected that that
//in TestCase.ExpectedCode.
func RunRequests(testCases []TestCase, t *testing.T) {
	for _, tc := range testCases {
		request(&tc, t)
	}
}

func request(tc *TestCase, t *testing.T) {
	req, err := http.NewRequest("GET", tc.Url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	Service.ServeHTTP(w, req)

	if w.Code != tc.ExpectedCode {
		t.Fatalf("Test [%s] failed: expected code %d, but got %d", tc.Description, tc.ExpectedCode, w.Code)
	}
}
