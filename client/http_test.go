package client

import (
	"fmt"
	"github.com/dooman87/kolibri/test"
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
	err, code := GetJson(ts.URL, s)

	test.Error(t,
		test.Nil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
		test.Equal(200, code, "code"),
		test.Equal("value", s.Field, "response"),
	)
}

func TestGetJsonErrorResponse(t *testing.T) {
	ts := startServer(404, "Not found")
	defer ts.Close()

	err, _ := GetJson(ts.URL, &S{})
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
	)
}

func TestGetJsonInvalidJson(t *testing.T) {
	ts := startServer(200, "That's not a json.")
	defer ts.Close()

	err, _ := GetJson(ts.URL, &S{})
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
	)
}

func TestGetJsonTargetIsNil(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err, _ := GetJson(ts.URL, nil)
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(0, httpCallsCount, "calls to http server"),
	)
}

func TestGetJsonUrlIsEmpty(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err, _ := GetJson("", &S{})
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(0, httpCallsCount, "calls to http server"),
	)
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
