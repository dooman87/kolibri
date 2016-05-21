package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type S struct {
	Field string
}

var (
	httpCallsCount = 0
)

func TestGetJson(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	s := &S{}
	err := GetJson(ts.URL, s)
	if err != nil {
		t.Fatal(err)
	}

	if s.Field != "value" {
		t.Fatalf("Expected field to be [%s], but got [%s]", "value", s.Field)
	}
	testHttpCallsCount(1, t)
}

func TestGetJsonErrorResponse(t *testing.T) {
	ts := startServer(404, "Not found")
	defer ts.Close()

	testError(GetJson(ts.URL, &S{}), t)
	testHttpCallsCount(1, t)
}

func TestGetJsonInvalidJson(t *testing.T) {
	ts := startServer(200, "That's not a json.")
	defer ts.Close()

	testError(GetJson(ts.URL, &S{}), t)
	testHttpCallsCount(1, t)
}

func TestGetJsonTargetIsNil(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err := GetJson(ts.URL, nil)
	testError(err, t)
	testHttpCallsCount(0, t)
}

func TestGetJsonUrlIsEmpty(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err := GetJson("", &S{})
	testError(err, t)
	testHttpCallsCount(0, t)
}

func testError(err error, t *testing.T) {
	if err == nil {
		t.Fatalf("Expected getting error")
	}
}

func testHttpCallsCount(expectedCallsCount int, t *testing.T) {
	if httpCallsCount != expectedCallsCount {
		t.Fatalf("Expected %d calls to http server but got %d", expectedCallsCount, httpCallsCount)
	}
}

func startServer(code int, resp string) *httptest.Server {
	httpCallsCount = 0
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpCallsCount++
		if code != 200 {
			w.WriteHeader(code)
		} else {
			fmt.Fprintln(w, resp)
		}
	}))
}
