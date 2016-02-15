package test_test
import (
	"github.com/dooman87/kolibri/test"
	"net/http"
	"testing"
)

func ExampleRunRequests() {
	test.Service = func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("OK"))
	}
	testCases := []test.TestCase{
		{"http://localhost:8080", http.StatusOK, "Should return 200",},
	}
	test.RunRequests(testCases, &testing.T{})
}