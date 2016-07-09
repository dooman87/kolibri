# Kolibri - not a framework for micro services.

[![Build Status](https://travis-ci.org/dooman87/kolibri.svg?branch=master)](https://travis-ci.org/dooman87/kolibri)

Kolibri is a set of lightweight helpers that I found useful to have 
during micro services development. Currently, it's providing next 
functionality:

* Http Client
* Health Check
* Testing services

One of the goal is to have minimum dependencies. Currently it requires
only Go 1.5+.

## Http Client

Getting JSON form HTTP endpoint and marshalling it to interface. 

## Health Checks

Just returns "OK". In future, I'd like to add set of checks, like 
http, mongo, etc.

## Service Testing

Provides TestCase abstraction that allow you to easily test your 
endpoints. Example:

```
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
```

## Contribution

Firstly, create a ticket and we can discuss a change. After that, you are welcome to create PRs and 
I'm happy to review and merge them. Happy Coding! 