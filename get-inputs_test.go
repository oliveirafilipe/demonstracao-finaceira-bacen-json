package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetInputsWindows(t *testing.T) {
	inMemory := strings.NewReader("123\r\n")

	inputs := getInputs(map[string]*Variable{
		"foo": {"foo", "Foo Description", "", `^\d{3}$`},
	}, inMemory)

	expectedInputs := map[string]string{"foo": "123"}

	if !reflect.DeepEqual(inputs, expectedInputs) {
		t.Fatalf("erro")
	}
}

func TestGetInputsUnix(t *testing.T) {
	inMemory := strings.NewReader("123\n")

	inputs := getInputs(map[string]*Variable{
		"foo": {"foo", "Foo Description", "", `^\d{3}$`},
	}, inMemory)

	expectedInputs := map[string]string{"foo": "123"}

	if !reflect.DeepEqual(inputs, expectedInputs) {
		t.Fatalf("erro")
	}
}

func TestShouldGetDefaultValue(t *testing.T) {
	inMemory := strings.NewReader("\n")

	inputs := getInputs(map[string]*Variable{
		"foo": {"foo", "Foo Description", "890", `^\d{3}$`},
	}, inMemory)

	expectedInputs := map[string]string{"foo": "890"}

	if !reflect.DeepEqual(inputs, expectedInputs) {
		t.Fatalf("erro")
	}
}
