package test

import (
	"testing"
)

type Assert interface {
	Exec(t *testing.T)
}

type equalAssert struct {
	Expected    interface{}
	Actual      interface{}
	Description string
}

//Creates new equal assert object.
func Equal(expected interface{}, actual interface{}, description string) Assert {
	return &equalAssert{
		Expected:    expected,
		Actual:      actual,
		Description: description,
	}
}

func (a *equalAssert) Exec(t *testing.T) {
	if a.Actual != a.Expected {
		t.Errorf("Expected [%v] %s but got [%v]", a.Expected, a.Description, a.Actual)
	}
}

type notEqualAssert struct {
	NotExpected interface{}
	Actual      interface{}
	Description string
}

//Creates new non-equal assert object.
func NotEqual(notExpected interface{}, actual interface{}, description string) Assert {
	return &notEqualAssert{
		NotExpected: notExpected,
		Actual:      actual,
		Description: description,
	}
}

func (a *notEqualAssert) Exec(t *testing.T) {
	if a.Actual == a.NotExpected {
		t.Errorf("Expected %s to be not equal to [%v]", a.Description, a.NotExpected)
	}
}

type nilAssert struct {
	Actual      interface{}
	Description string
}

func (a *nilAssert) Exec(t *testing.T) {
	if a.Actual != nil {
		t.Errorf("Expected %s to be nil", a.Description)
	}
}

//Creates new nil assert object.
func Nil(actual interface{}, description string) Assert {
	return &nilAssert{
		Actual:      actual,
		Description: description,
	}
}

type notNilAssert struct {
	Actual      interface{}
	Description string
}

func (a *notNilAssert) Exec(t *testing.T) {
	if a.Actual == nil {
		t.Errorf("Expected %s to not be nil", a.Description)
	}
}

//Creates new not nil assert object.
func NotNil(actual interface{}, description string) Assert {
	return &notNilAssert{
		Actual:      actual,
		Description: description,
	}
}

//Test each assert and error on each failed. Using only plain == to compare expected
// and actual values.
// Error message has format: "Expected [%expected] %description but got [%actual]".
func Error(t *testing.T, a ...Assert) {
	for _, assert := range a {
		assert.Exec(t)
	}
}
