package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func TestShouldGenerateJSONStatement(t *testing.T) {
	input := [][]string{
		{"", "BAR", "FOO"},
		{"Value1", "100", "200"},
		{""},
		{"Value2", "50.123", "100.123"},
		{" Value2.1", "30.123", "70"},
		{" Value2.2", "20", "30"},
		{"  Value2.2.1", "20", "30.123"},
		{"Value3"},
	}
	inputBaseDates := map[string]string{"BAR": "dt1", "FOO": "dt2"}
	output, _ := json.MarshalIndent(processStatements(input, inputBaseDates), "", "  ")

	content, err := ioutil.ReadFile("testdata/generated_statements.golden.json")
	if err != nil {
		t.Fatalf("Error loading golden file: %s", err)
	}

	got := strings.TrimSpace(string(output))
	expectedOutput := strings.TrimSpace(string(content))
	if got != expectedOutput {
		t.Errorf("Want:\n%s\nGot:\n%s", expectedOutput, got)
	}
}
