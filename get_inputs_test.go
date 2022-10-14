package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
	"warrenbrasil/demonstracao-finaceira-bacen-json/inputvar"
)

func TestGetInputsWindows(t *testing.T) {
	inMemory := strings.NewReader("123\r\n")

	var inputObj, err = inputvar.Create(inputvar.Options{
		Message:    "Foo Description",
		Default:    "",
		Validation: *regexp.MustCompile(`^\d{3}$`),
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	GetInputs(map[string]*inputvar.Input{
		"foo": inputObj,
	}, inMemory)

	expectedValue := "123"

	if !reflect.DeepEqual(inputObj.Value, expectedValue) {
		t.Fatalf("Got: %s\nWant: %s", inputObj.Value, expectedValue)
	}
}

func TestGetInputsUnix(t *testing.T) {
	inMemory := strings.NewReader("123\n")

	var inputObj, err = inputvar.Create(inputvar.Options{
		Message:    "Foo Description",
		Default:    "",
		Validation: *regexp.MustCompile(`^\d{3}$`),
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	GetInputs(map[string]*inputvar.Input{
		"foo": inputObj,
	}, inMemory)

	expectedValue := "123"

	if !reflect.DeepEqual(inputObj.Value, expectedValue) {
		t.Fatalf("Got: %s\nWant: %s", inputObj.Value, expectedValue)
	}
}

func TestShouldGetDefaultValue(t *testing.T) {
	inMemory := strings.NewReader("\n")

	var inputObj, err = inputvar.Create(inputvar.Options{
		Message:    "Foo Description",
		Default:    "890",
		Validation: *regexp.MustCompile(`^\d{3}$`),
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	GetInputs(map[string]*inputvar.Input{
		"foo": inputObj,
	}, inMemory)

	expectedValue := "890"

	if !reflect.DeepEqual(inputObj.Value, expectedValue) {
		t.Fatalf("Got: %s\nWant: %s", inputObj.Value, expectedValue)
	}
}
