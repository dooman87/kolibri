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

Getting/sending JSON to HTTP endpoint and marshalling it to interface.
Supporting GET and POST methods.

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

## Testing helpers

* Test multiple statements with asserts. Simplifies code like

```
    if data.Id != 1 {
        t.Errorf("Expected [%d] id but got [%d]", 1, data.Id)
    }
    if data.Squirell == nil {
        t.Errorf("Expected squirell to not be nil", "Jack", data.SquirellName)
    }
```

to

```
	test.Error(t,
		test.Equal(1, data.Id, "id"),
		test.NotNil(data.SquirellName, "squirell"),
	)
```

## Contribution

Firstly, create a ticket and we can discuss a change. After that, 
you are more then welcome to create PRs and I'm happy to review and 
merge them. Happy Coding! 