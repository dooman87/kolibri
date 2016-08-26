//Provides helper to test http handlers.
package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
)

type ResponseHandler func(w *httptest.ResponseRecorder, t *testing.T)

//Describes test case. Each test case
//will be executed as http request against
//Service handler.
type TestCase struct {
	//You can use this field in case you need simple GET request
	//or pass request using Request field.
	Url          string
	//Request to execute. If nil then Url will be used and GET request will be executed.
	Request      *http.Request
	//Expected respose code. If not passed then will check for 200.
	ExpectedCode int
	Description  string
	//Optional handler that will be called at the end of the test case.
	Handler      ResponseHandler
}

var (
	//Handler that will be executed for test cases.
	Service http.HandlerFunc
	//Test context
	T *testing.T
)

//Runs all test cases and failing test if
//response code is not equals to expected that that
//in TestCase.ExpectedCode.
func RunRequests(testCases []TestCase) {
	for _, tc := range testCases {
		request(&tc, T)
	}
}

// Wraps http.NewRequest to process error case.
// If http.NewRequest returns error then executing interrupts execution with T.Fatal()
func NewRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		T.Fatal(err)
	}
	return req
}

func request(tc *TestCase, t *testing.T) {
	var (
		req *http.Request
		err error
	)

	if len(tc.Url) > 0 {
		req, err = http.NewRequest("GET", tc.Url, nil)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		req = tc.Request
	}

	w := httptest.NewRecorder()

	Service.ServeHTTP(w, req)

	if (tc.ExpectedCode == 0 && w.Code != http.StatusOK) || w.Code != tc.ExpectedCode {
		t.Fatalf("Test [%s] failed: expected code %d, but got %d", tc.Description, tc.ExpectedCode, w.Code)
	}

	if tc.Handler != nil {
		tc.Handler(w, t)
	}
}
