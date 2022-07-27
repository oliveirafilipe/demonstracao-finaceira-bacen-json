package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	financialStatements := FinancialStatements{
		Cnpj:             "foo-cnpj",
		DocumentCode:     "foo-doc-code",
		TypeRemittance:   "foo-type",
		ValuesMultiplier: 0,
		BaseDate:         "foo-base-date",
		BaseDatesReferences: []BaseDatesReference{
			{
				Id:   "dt1",
				Date: "FooDate1",
			},
			{
				Id:   "dt2",
				Date: "FooDate2",
			},
		},
		BalancoPatrimonial: BalancoT{},
		DRE:                DRET{},
		Caixa:              CaixaT{},
		DMPL:               DMPLT{},
		DRA:                DRAT{},
	}

	// TODO: Na real nao preciso ler os CSVs do folder ja que ja sei os nomes deles

	for _, file := range getCSVsInFolder() {
		lines := openCsv(file)
		var statements []Statement = processStatemets(lines, file)

		if strings.Contains(file, "caixa") {
			financialStatements.Caixa = CaixaT{
				Statements: statements,
			}
		} else if strings.Contains(file, "balanco") {
			financialStatements.BalancoPatrimonial = BalancoT{
				Statements: statements,
			}
		} else if strings.Contains(file, "dmpl") {
			financialStatements.DMPL = DMPLT{
				Statements: statements,
			}
		} else if strings.Contains(file, "dra") {
			financialStatements.DRA = DRAT{
				Statements: statements,
			}
		} else if strings.Contains(file, "dre") {
			financialStatements.DRE = DRET{
				Statements: statements,
			}
		}
	}

	data, _ := json.Marshal(financialStatements)
	fmt.Println(string(data))

	var foo string

	fmt.Printf("Aperte ENTER para finalizar...")
	fmt.Scanln(&foo)

}

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

		//fmt.Println(line[0], levelFoo, arrayIndex, statements, individualizedValues)

		statementId++
	}

	return statements
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

/*
func main() {

	pigeon := FinancialStatements{
		Cnpj: "92875780",
		DocumentCode: "9011",
		TypeRemittance: "I",
		ValuesMultiplier: 1000,
		BaseDate: "062020",
		BaseDatesReferences: []BaseDatesReference{
			{
				Id: "dt1",
				Date: "A122021",
			},
			{
				Id: "dt2",
				Date: "A122020",
			},
		},
		BalancoPatrimonial: BalancoT{
			Statements: []Statement{
				{
					Id: "conta1",
					Description: "RECEITAS DA INTERMEDIAÇÃO FINANCEIRA",
					Level: "1",
					ParentStatement: "",
					IndividualizedValues: []IndividualizedValue{
						{
							BaseDate: "dt1",
							Ammount: 4,
						},
						{
							BaseDate: "dt2",
							Ammount: 5992,
						},
					},
				},
			},
		},
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))
}*/
