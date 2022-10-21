package inputvar

import (
	"fmt"
	"regexp"
	"testing"
)

func TestErrorForInvalidDefault(t *testing.T) {

	var _, err = New(Options{
		Message:    "Foo",
		Default:    "foo",
		Validation: *regexp.MustCompile("bar"),
	})

	if err == nil {
		t.Fatalf("expected error to not be nil")
	}
}

func TestValueShouldBeEqualToDefault(t *testing.T) {

	var inputObj, err = New(Options{
		Message:    "Foo",
		Default:    "123",
		Validation: *regexp.MustCompile(`^\d{3}$`),
	})

	if err != nil {
		t.Errorf("expected error to be nil")
	}

	if inputObj.Value != inputObj.Default {
		t.Errorf("Expected %s; Got: %s", inputObj.Default, inputObj.Value)
	}
}

func TestShouldGetErrorForInvalidValue(t *testing.T) {

	var inputObj, err = New(Options{
		Message:    "Foo",
		Default:    "123",
		Validation: *regexp.MustCompile(`^\d{3}$`),
	})

	if err != nil {
		t.Errorf("expected error to be nil")
	}

	var err2 = inputObj.SetValue("abc")

	fmt.Println(err2)

	if err2 == nil {
		t.Errorf("expected error to be nil")
	}
}
