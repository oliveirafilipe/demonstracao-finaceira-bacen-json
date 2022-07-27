package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func openCsv(fileName string) [][]string {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvr := csv.NewReader(csvFile)
	csvr.FieldsPerRecord = -1
	csvLines, err := csvr.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return csvLines
}
