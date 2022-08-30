package main

import "fmt"

func getBaseDates(files []string) map[string]string {
	baseDatesMap := map[string]string{}
	i := 1
	for _, file := range files {
		lines := openCsv(file)
		for _, line := range lines {
			for _, date := range line {
				//https://stackoverflow.com/a/2050629/13152732
				if _, ok := baseDatesMap[date]; !ok && date != "" {
					baseDatesMap[date] = fmt.Sprint("dt", i)
					i++
				}
			}
			break
		}

	}

	return baseDatesMap
}

// statement := Statement{
// 	Id:                   "conta" + strconv.Itoa(statementId),
// 	Description:          strings.TrimSpace(line[0]),
// 	Level:                arrayToString(level, "."),
// 	ParentStatement:      "",
// 	IndividualizedValues: individualizedValues,
// }

// if len(parentStatements) >= 2 && len(parentStatements[len(parentStatements)-2]) != 0 {
// 	statement.ParentStatement = fmt.Sprint("conta", parentStatements[len(parentStatements)-2])
// }

// statements = append(statements, statement)
