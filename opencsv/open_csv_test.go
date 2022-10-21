package opencsv

import (
	"reflect"
	"testing"
)

func TestShouldReadFileCorrectly(t *testing.T) {
	expectedContent := [][]string{{"", "", "", "BAR", "FOO", "ZEDD"}, {"a", "b", "c", "1", "2", "3"}}
	content, err := OpenCsv("testdata/test_input1.csv")

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(content, expectedContent) {
		t.Fatalf(`openCsv("testdata/test_input1.csv") = %s, expected %s`, content, expectedContent)
	}
}
