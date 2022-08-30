package main

import (
	"fmt"
	"strconv"
	"strings"
)

var inputDefinitions = map[string]*Variable{
	"cnpj":       {"CNPJ", "Raiz do CNPJ. Numero de 8 dígitos", "92875780", `^\d{8}$`},
	"doccode":    {"Código do Documento", "Código do Documento. Consulte documentação oficial.", "9011", `^9\d{2}1$`},
	"type":       {"Tipo de  Remeça", "I (Inclusão) ou S (Substituição)", "I", `^(I|S)$`},
	"multiplier": {"Unidade de Medida", "Indica o multiplicado adotado para os valores em reais ", "1000", `^\d+$`},
	"basedate":   {"Data Base", "Mês (2 dígitos) + Ano (4 dígitos) - Referente ao último mês do período analisado", "", `^\d{6}$`},
}

func main() {
	var inputs = getInputs(inputDefinitions)
	var sourceFiles = []string{"balanco.csv", "caixa.csv", "dmpl.csv", "dra.csv", "dre.csv"}

	multiplier, _ := strconv.ParseInt(inputs["multiplier"], 10, 32)
	financialStatements := FinancialStatements{
		Cnpj:             inputs["cnpj"],
		DocumentCode:     inputs["doccode"],
		TypeRemittance:   inputs["type"],
		ValuesMultiplier: int(multiplier),
		BaseDate:         inputs["basedate"],
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
		var statements []Statement = processStatemets(lines)

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
