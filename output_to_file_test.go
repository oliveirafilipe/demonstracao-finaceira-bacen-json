package main

import (
	"os"
	"testing"
)

func TestShouldSaveToFile(t *testing.T) {
	type Foo struct {
		Foo  int
		Bar  []int
		Zedd string
	}
	var foo = Foo{Foo: 123, Bar: []int{1, 2, 3}, Zedd: "abc"}
	outputToFile(foo, "testdata/output.json")

	fileContent, err := os.ReadFile("testdata/output.json")
	expectedFileContent := `{
  "Foo": 123,
  "Bar": [
    1,
    2,
    3
  ],
  "Zedd": "abc"
}`

	if err != nil {
		t.Fatalf("Error while opening file; %s", err)
	}

	if string(fileContent) != expectedFileContent {
		t.Fatalf("File content differs from expected;\n Received: %s\n Expected: %s", string(fileContent), expectedFileContent)
	}
}
