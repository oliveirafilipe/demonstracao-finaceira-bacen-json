package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

func getCSVsInFolder(base string) []string {
	var CSVs []string
	files, err := ioutil.ReadDir(base)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			match, _ := regexp.MatchString("\\.csv$", file.Name())
			if match {
				CSVs = append(CSVs, file.Name())
			}

		}

	}

	return CSVs
}
