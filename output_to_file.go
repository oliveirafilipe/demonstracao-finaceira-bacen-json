package main

import (
	"encoding/json"
	"io/ioutil"
)

func outputToFile(data interface{}, fileName string) error {
	file, _ := json.MarshalIndent(data, "", "  ")

	return ioutil.WriteFile(fileName, file, 0644)
}
