package datasource

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"warrenbrasil/demonstracao-finaceira-bacen-json/finstm"
)

type CSV struct {
	Path string
}

func CheckFile(path string) error {
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
	}
	return err
}

func openCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvr := csv.NewReader(file)
	csvr.FieldsPerRecord = -1
	return csvr.ReadAll()
}

func (csv *CSV) GetBaseDates() ([]string, error) {
	baseDates := []string{}

	lines, err := openCsv(csv.Path)

	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil,  errors.New("Empty CSV")
	}

	for _, column := range lines[0] {
		if column != "" {
			baseDates = append(baseDates, column)
		}
	}

	return baseDates, nil
}

func (csv *CSV) GetStatements(baseDatesMap map[string]string) ([]finstm.Statement, error) {
	lines, err := openCsv(csv.Path)

	if err != nil {
		return nil, err
	}

	return ProcessStatements(lines, baseDatesMap), nil

}

func ProcessStatements(lines [][]string, baseDates map[string]string) []finstm.Statement {
	statementId := 1
	var statements []finstm.Statement = []finstm.Statement{}
	var level []int = []int{}
	var parentStatements []string = []string{}
	for _, line := range lines[1:] {
		match, _ := regexp.MatchString("^(-|\\s)*$", line[0])
		if match {
			continue
		}

		var individualizedValues []finstm.IndividualizedValue = []finstm.IndividualizedValue{}

		for i := 1; i < len(line); i++ {
			match, _ := regexp.MatchString(`^(-|\s)*$`, line[i])
			if !match {
				amount, _ := strconv.ParseFloat(line[i], 64)
				individualizedValues = append(individualizedValues, finstm.IndividualizedValue{
					BaseDate: baseDates[lines[0][i]],
					Amount:   amount,
				})
			}
		}

		if len(individualizedValues) == 0 {
			continue
		}

		levelIdx := len(regexp.MustCompile(`^\s*`).FindString(line[0]))
		if len(level) < levelIdx+1 {
			level = append(level, 0)
			parentStatements = append(parentStatements, strconv.Itoa(statementId))
		} else {
			parentStatements = parentStatements[:levelIdx+1]
			level = level[:levelIdx+1]
		}
		level[levelIdx]++
		parentStatements[len(parentStatements)-1] = strconv.Itoa(statementId)

		statement := finstm.Statement{
			Id:                   "conta" + strconv.Itoa(statementId),
			Description:          strings.TrimSpace(line[0]),
			Level:                arrayToString(level, "."),
			ParentStatement:      "",
			IndividualizedValues: individualizedValues,
		}

		if len(parentStatements) >= 2 && len(parentStatements[len(parentStatements)-2]) != 0 {
			statement.ParentStatement = fmt.Sprint("conta", parentStatements[len(parentStatements)-2])
		}

		statements = append(statements, statement)

		statementId++
	}

	return statements
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
