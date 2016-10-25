package client_test

import (
	"fmt"
	"github.com/dooman87/kolibri/client"
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
	err, code := (&client.Http{}).GetJson(ts.URL, s)

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

	err, _ := (&client.Http{}).GetJson(ts.URL, &S{})
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
	)
}

func TestGetJsonInvalidJson(t *testing.T) {
	ts := startServer(200, "That's not a json.")
	defer ts.Close()

	err, _ := (&client.Http{}).GetJson(ts.URL, &S{})
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
	)
}

func TestGetJsonTargetIsNil(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err, _ := (&client.Http{}).GetJson(ts.URL, nil)
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(0, httpCallsCount, "calls to http server"),
	)
}

func TestGetJsonUrlIsEmpty(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err, _ := (&client.Http{}).GetJson("", &S{})
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(0, httpCallsCount, "calls to http server"),
	)
}

func TestPostJson(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	s := &S{
		Field: "value",
	}
	resp := &S{}
	err, _ := (&client.Http{}).PostJson(ts.URL, s, resp)
	test.Error(t,
		test.Nil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
		test.Equal("value", resp.Field, "response"),
	)
}

func TestPostJsonUrlIsEmpty(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	err, _ := (&client.Http{}).PostJson("", nil, nil)
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(0, httpCallsCount, "calls to http server"),
	)
}

func TestPostJsonBodyIsNil(t *testing.T) {
	ts := startServer(200, `{"field": "value"}`)
	defer ts.Close()

	resp := &S{}
	err, _ := (&client.Http{}).PostJson(ts.URL, nil, resp)
	test.Error(t,
		test.Nil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
		test.Equal("value", resp.Field, "response"),
	)
}

func TestPostJsonResponseIsNil(t *testing.T) {
	ts := startServer(200, ``)
	defer ts.Close()

	s := &S{
		Field: "value",
	}
	err, _ := (&client.Http{}).PostJson(ts.URL, s, nil)
	test.Error(t,
		test.Nil(err, "error"),
		test.Equal(1, httpCallsCount, "calls to http server"),
	)
}

func TestPostJsonErrorResponse(t *testing.T) {
	ts := startServer(500, `Error response`)
	defer ts.Close()

	resp := &S{}
	s := &S{
		Field: "value",
	}
	err, code := (&client.Http{}).PostJson(ts.URL, s, resp)
	test.Error(t,
		test.NotNil(err, "error"),
		test.Equal(500, code, "response status code"),
		test.Equal(1, httpCallsCount, "calls to http server"),
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
