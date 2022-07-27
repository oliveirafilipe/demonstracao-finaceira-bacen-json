package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var sourceFiles = []string{"balanco.csv", "caixa.csv", "dmpl.csv", "dra.csv", "dre.csv"}

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

	missingFiles := checkRequiredFiles(sourceFiles)
	if len(missingFiles) > 0 {
		for _, missingFile := range checkRequiredFiles(sourceFiles) {
			fmt.Println(fmt.Sprint("- Arquivo obrigatório não encontrado ", missingFile))

		}
		enterToClose()
		return
	}

	for _, file := range sourceFiles {
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
	_ = data
	fmt.Println(string(data)[:400])

}
