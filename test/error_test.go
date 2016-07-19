package test_test

import (
	"errors"
	"github.com/dooman87/kolibri/test"
	"testing"
)

func TestError(t *testing.T) {
	test.Error(t,
		test.Equal(1, 1, "apples"),
		test.NotEqual(1, 2, "oranges"),
		test.NotNil("Bob", "student name"),
		test.Nil(nil, "error"),
	)
}

func ExampleError() {
	var t = new(testing.T)
	test.Error(t,
		//Error message will be: Expected [1] apples but got [6]
		test.Equal(1, 6, "apples"),
		//Error message will be: Expected price to be not equal to 2.99
		test.NotEqual(2.99, 2.99, "price"),
		//Error message will be: Expected error to be nil
		test.Nil(errors.New("ERROR"), "error"),
		//Error message will be: Expected student name to not be nil
		test.NotNil(nil, "student name"),
	)
}
