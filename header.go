package main

import "github.com/oliveirafilipe/demonstracao-finaceira-bacen-json/metadata"

func getHeader() string {

	var header string = "Demonstrativos Financeiros BACEN - Converte CSV para JSON\n" +
		"(Versão: " + metadata.Version + " - Autor: @oliveirafilipe)\n\n" +
		"Para instruções de uso, ou novas versões, acesse: https://github.com/oliveirafilipe/demonstracao-finaceira-bacen-json\n"

	if metadata.Version == "development" {
		header = header + "==> Você provavelmente não deveria estar usando essa versão.\n"
	}

	return header
}
