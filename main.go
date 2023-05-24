package main

import (
	"fmt"
	"os"
	"strconv"
	"warrenbrasil/demonstracao-finaceira-bacen-json/basedates"
	"warrenbrasil/demonstracao-finaceira-bacen-json/datasource"
	"warrenbrasil/demonstracao-finaceira-bacen-json/finstm"
	"warrenbrasil/demonstracao-finaceira-bacen-json/input"
)

type Source struct {
	Path       string
	DataSource *datasource.CSV
	Statements []finstm.Statement
}

func main() {
	fmt.Println(header)

	var sourceFiles = map[string]*Source{
		"balanco": {
			Path: "./balanco.csv",
		},
		"caixa": {
			Path: "./caixa.csv",
		},
		"dmpl": {
			Path: "./caixa.csv",
		},
		"dra": {
			Path: "./dra.csv",
		},
		"dre": {
			Path: "./dre.csv",
		},
	}
	var errorFlag = false

	for _, el := range sourceFiles {
		err := datasource.CheckFile(el.Path)
		if err != nil {
			errorFlag = true
			if err == err.(*os.PathError) {
				fmt.Println(fmt.Sprint("- Arquivo obrigatório não encontrado ", el.Path))
			} else {
				fmt.Print("ERRO: Falha ao verificar existências dos arquivos.")
			}
		} else {
			el.DataSource = &datasource.CSV{Path: el.Path}
		}
	}
	if errorFlag {
		fmt.Println("DICA: Se você está no Windows, talvez você precise retirar a extensão (.csv) do nome do arquivo.")
		input.EnterToClose()
		return
	}

	var questions = Questions()
	input.GetInputs(questions, os.Stdin)

	// // ====== GET DATES FROM FILES ======
	dates := []string{}
	for _, el := range sourceFiles {
		baseDates, err := el.DataSource.GetBaseDates()
		if err != nil {
			fmt.Print("ERRO: Erro na busca de datas. Verifique se todo os arquivos possuem conteudo. Encerrando...")
			input.EnterToClose()
			return
		} else {
			dates = append(dates, baseDates...)
		}
	}

	// // ====== GENERATE IDs FOR DATES ======
	var baseDatesMap = basedates.GenerateIDsForDates(dates)
	var baseDates []finstm.BaseDatesReference
	for date, id := range baseDatesMap {
		baseDatesReference := finstm.BaseDatesReference{
			Id:   id,
			Date: date,
		}
		baseDates = append(baseDates, baseDatesReference)
	}

	// // ====== GET STATEMENTS FROM FILES, USING DATES AND THEIR IDs ======
	for _, file := range sourceFiles {
		statements, err := file.DataSource.GetStatements(baseDatesMap)
		if err != nil {
			fmt.Print("ERRO: Falha ao buscar demonstrativos do arquivo.")
			input.EnterToClose()
			return
		}
		file.Statements = statements
	}

	multiplier, _ := strconv.ParseInt(questions["multiplier"].Value, 10, 32)
	financialStatements := finstm.FinancialStatements{
		Cnpj:                questions["cnpj"].Value,
		DocumentCode:        questions["doccode"].Value,
		TypeRemittance:      questions["type"].Value,
		ValuesMultiplier:    int(multiplier),
		BaseDate:            questions["basedate"].Value,
		BaseDatesReferences: baseDates,
		BalancoPatrimonial: finstm.BalancoT{
			Statements: sourceFiles["balanco"].Statements,
		},
		DRE: finstm.DRET{
			Statements: sourceFiles["dre"].Statements,
		},
		Caixa: finstm.CaixaT{
			Statements: sourceFiles["caixa"].Statements,
		},
		DMPL: finstm.DMPLT{
			Statements: sourceFiles["dmpl"].Statements,
		},
		DRA: finstm.DRAT{
			Statements: sourceFiles["dra"].Statements,
		},
	}

	if err := financialStatements.Save(); err == nil {
		fmt.Print("Arquivo de saída (resultado.json) gerado com sucesso!")
	} else {
		fmt.Print("ERRO: Falha ao gerar arquivo de saída. Execute o programa novamente!")
	}

	input.EnterToClose()
}
