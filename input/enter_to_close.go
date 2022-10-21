package input

import "fmt"

func EnterToClose() {
	var foo string

	fmt.Printf("\nAperte ENTER para finalizar...")
	fmt.Scanln(&foo)
}
