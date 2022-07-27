package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func processStatemets(lines [][]string, name string) []Statement {
	_ = name
	statementId := 1
	var statements []Statement = []Statement{}
	var level []int = []int{}
	var parentStatements []string = []string{}
	for _, line := range lines[1:] {
		match, _ := regexp.MatchString("^(-|\\s)*$", line[0])
		if match {
			continue
		}

		// arrayIndex := statementId -1
		levelFoo := len(regexp.MustCompile("^\\s*").FindString(line[0]))
		if len(level) < levelFoo+1 {
			level = append(level, 0)
			parentStatements = append(parentStatements, strconv.Itoa(statementId))
		} else {
			parentStatements = parentStatements[:levelFoo+1]
			level = level[:levelFoo+1]
		}
		level[levelFoo]++
		parentStatements[len(parentStatements)-1] = strconv.Itoa(statementId)

		var individualizedValues []IndividualizedValue = []IndividualizedValue{}

		for i := 1; i < len(line); i++ {
			match, _ := regexp.MatchString("^(-|\\s)*$", line[i])
			if !match {
				ammount, _ := strconv.ParseFloat(line[i], 64)
				individualizedValues = append(individualizedValues, IndividualizedValue{
					BaseDate: fmt.Sprint("dt", i),
					Ammount:  ammount,
				})
			}
		}

		if len(individualizedValues) == 0 {
			continue
		}

		statement := Statement{
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
