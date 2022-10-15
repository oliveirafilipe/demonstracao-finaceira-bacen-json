package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"warrenbrasil/demonstracao-finaceira-bacen-json/finstm"
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
	var baseDates []finstm.BaseDatesReference
	for date, id := range baseDatesMap {
		baseDatesReference := finstm.BaseDatesReference{
			Id:   id,
			Date: date,
		}
		baseDates = append(baseDates, baseDatesReference)
	}

	multiplier, _ := strconv.ParseInt(inputs["multiplier"].Value, 10, 32)
	financialStatements := finstm.FinancialStatements{
		Cnpj:                inputs["cnpj"].Value,
		DocumentCode:        inputs["doccode"].Value,
		TypeRemittance:      inputs["type"].Value,
		ValuesMultiplier:    int(multiplier),
		BaseDate:            inputs["basedate"].Value,
		BaseDatesReferences: baseDates,
		BalancoPatrimonial:  finstm.BalancoT{},
		DRE:                 finstm.DRET{},
		Caixa:               finstm.CaixaT{},
		DMPL:                finstm.DMPLT{},
		DRA:                 finstm.DRAT{},
	}

	for _, file := range sourceFiles {
		lines := openCsv(file)
		var statements []finstm.Statement = processStatements(lines, baseDatesMap)

		if strings.Contains(file, "caixa") {
			financialStatements.Caixa = finstm.CaixaT{
				Statements: statements,
			}
		} else if strings.Contains(file, "balanco") {
			financialStatements.BalancoPatrimonial = finstm.BalancoT{
				Statements: statements,
			}
		} else if strings.Contains(file, "dmpl") {
			financialStatements.DMPL = finstm.DMPLT{
				Statements: statements,
			}
		} else if strings.Contains(file, "dra") {
			financialStatements.DRA = finstm.DRAT{
				Statements: statements,
			}
		} else if strings.Contains(file, "dre") {
			financialStatements.DRE = finstm.DRET{
				Statements: statements,
			}
		}
	}

	if err := financialStatements.Save(); err == nil {
		fmt.Print("Arquivo de saída (resultado.json) gerado com sucesso!")
	} else {
		fmt.Print("Falha ao gerar arquivo de saída. Execute o programa novamente!")
	}

	enterToClose()
}
