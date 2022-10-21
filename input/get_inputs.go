package input

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"warrenbrasil/demonstracao-finaceira-bacen-json/inputvar"
)

type Variable struct {
	Name        string
	Description string
	Default     string
	Validation  string
}

func GetInputs(vars map[string]*inputvar.Input, reader io.Reader) error {

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
			fmt.Printf("\nVariable: %s\n", variable.Message)

			if variable.Default != "" {
				fmt.Printf("   ENTER para valor Default: %s\n", variable.Default)
			}
			reader := bufio.NewReader(reader)
			line, err := reader.ReadString('\n')

			if err != nil {
				return err
			}

			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")

			if line == "" {
				if variable.IsEmpty() {
					fmt.Println("Um valor deve ser informado")
					continue
				}
			} else {
				if err := variable.SetValue(line); err != nil {
					fmt.Println("Valor informado não é valido")
					continue
				}
			}
			break
		}
	}

	return nil
}
