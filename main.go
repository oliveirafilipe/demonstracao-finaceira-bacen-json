package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(header)

	var sourceFiles = []string{"balanco.csv", "caixa.csv", "dmpl.csv", "dra.csv", "dre.csv"}

	missingFiles := checkRequiredFiles(sourceFiles, ".")
	if len(missingFiles) > 0 {
		for _, missingFile := range missingFiles {
			fmt.Println(fmt.Sprint("- Arquivo obrigatório não encontrado ", missingFile))
		}
		enterToClose()
		return
	}

	var inputs = Questions()
	GetInputs(inputs, os.Stdin)

	var baseDatesMap = getBaseDates(sourceFiles)
	var baseDates []BaseDatesReference
	for date, id := range baseDatesMap {
		baseDatesReference := BaseDatesReference{
			Id:   id,
			Date: date,
		}
		baseDates = append(baseDates, baseDatesReference)
	}

	multiplier, _ := strconv.ParseInt(inputs["multiplier"].Value, 10, 32)
	financialStatements := FinancialStatements{
		Cnpj:                inputs["cnpj"].Value,
		DocumentCode:        inputs["doccode"].Value,
		TypeRemittance:      inputs["type"].Value,
		ValuesMultiplier:    int(multiplier),
		BaseDate:            inputs["basedate"].Value,
		BaseDatesReferences: baseDates,
		BalancoPatrimonial:  BalancoT{},
		DRE:                 DRET{},
		Caixa:               CaixaT{},
		DMPL:                DMPLT{},
		DRA:                 DRAT{},
	}

	for _, file := range sourceFiles {
		lines := openCsv(file)
		var statements []Statement = processStatements(lines, baseDatesMap)

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

	if outputToFile(financialStatements, "resultado.json") == nil {
		fmt.Print("Arquivo de saída (resultado.json) gerado com sucesso!")
	} else {
		fmt.Print("Falha ao gerar arquivo de saída. Execute o programa novamente!")
	}

	enterToClose()
}
