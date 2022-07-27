package main

import (
	"strings"
)

func checkRequiredFiles(requiredFiles []string) []string {
	files := getCSVsInFolder()
	joinedFiles := strings.Join(files, ";")
	var missingFiles []string = []string{}

	for _, requiredFile := range requiredFiles {
		if !strings.Contains(joinedFiles, requiredFile) {
			missingFiles = append(missingFiles, requiredFile)
		}
	}

	return missingFiles
}
