package opencsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func OpenCsv(fileName string) ([][]string, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvr := csv.NewReader(csvFile)
	csvr.FieldsPerRecord = -1
	return csvr.ReadAll()
}
