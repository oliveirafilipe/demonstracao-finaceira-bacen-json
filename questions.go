package main

import (
	"regexp"

	"github.com/oliveirafilipe/demonstracao-finaceira-bacen-json/inputvar"
)

func Questions() map[string]*inputvar.Input {
	var inputDefinitions = make(map[string]*inputvar.Input)
	var inputObj, _ = inputvar.New(inputvar.Options{
		Message:    "Raiz do CNPJ. Numero de 8 dígitos",
		Default:    "92875780",
		Validation: *regexp.MustCompile(`^\d{8}$`),
	})
	inputDefinitions["cnpj"] = inputObj
	inputObj, _ = inputvar.New(inputvar.Options{
		Message:    "Mês (2 dígitos) + Ano (4 dígitos) - Referente ao último mês do período analisado",
		Default:    "",
		Validation: *regexp.MustCompile(`^\d{6}$`),
	})
	inputDefinitions["basedate"] = inputObj
	inputObj, _ = inputvar.New(inputvar.Options{
		Message:    "Código do Documento. Consulte documentação oficial.",
		Default:    "9011",
		Validation: *regexp.MustCompile(`^9\d{2}1$`),
	})
	inputDefinitions["doccode"] = inputObj
	inputObj, _ = inputvar.New(inputvar.Options{
		Message:    "Tipo de Remeça. I (Inclusão) ou S (Substituição)",
		Default:    "I",
		Validation: *regexp.MustCompile(`^(I|S)$`),
	})
	inputDefinitions["type"] = inputObj
	inputObj, _ = inputvar.New(inputvar.Options{
		Message:    "Unidade de Medida. Indica o multiplicado adotado para os valores em reais",
		Default:    "1000",
		Validation: *regexp.MustCompile(`^\d+$`),
	})
	inputDefinitions["multiplier"] = inputObj
	return inputDefinitions
}
