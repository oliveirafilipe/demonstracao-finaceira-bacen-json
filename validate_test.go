package main

import (
	"reflect"
	"testing"
)

func TestShouldValidateFiles(t *testing.T) {
	missingFiles := checkRequiredFiles([]string{"test_input1.csv", "test_input2.csv"}, "testdata")
	if len(missingFiles) != 0 {
		t.Fatalf("checkRequiredFiles, missing following files %s", missingFiles)
	}
}

func TestShouldIndicateMissingFiles(t *testing.T) {
	missingFiles := checkRequiredFiles([]string{"foo.csv", "bar.csv"}, "testdata")
	expectedMissingFiles := []string{"foo.csv", "bar.csv"}
	if !reflect.DeepEqual(missingFiles, expectedMissingFiles) {
		t.Fatalf("checkRequiredFiles. Got: %s; Expected: %s", missingFiles, expectedMissingFiles)
	}
}
