package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

type Variable struct {
	Name        string
	Description string
	Default     string
	Validation  string
}

func getInputs(vars map[string]*Variable, reader io.Reader) map[string]string {
	var line string
	var err error

	var result = make(map[string]string)

	keys := make([]string, len(vars))

	i := 0
	for k := range vars {
		keys[i] = k
		i++
	}

	sort.Strings(keys) // prompt in lexical order

	for _, key := range keys {
		variable := vars[key]
		for {
			fmt.Printf("\nVariable: %s\n", variable.Name)
			fmt.Printf("   Descrição: %s\n", variable.Description)

			if variable.Default != "" {
				fmt.Printf("   ENTER para valor Default: %s\n", variable.Default)
			}
			bufa := bufio.NewReader(reader)
			line, err = bufa.ReadString('\n')

			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")

			if line == "" {
				if variable.Default == "" {
					fmt.Println("Um valor deve ser informado")
					continue
				} else {
					line = variable.Default
				}
			} else if variable.Validation != "" {
				r := regexp.MustCompile(variable.Validation)
				if !r.Match([]byte(line)) {
					fmt.Println("Valor informado não é valido")
					continue
				}

			}
			break
		}
		_ = err
		result[key] = line
	}

	return result
}
