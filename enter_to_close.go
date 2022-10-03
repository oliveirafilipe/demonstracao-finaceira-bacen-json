package main

import "fmt"

func enterToClose() {
	var foo string

	fmt.Printf("\nAperte ENTER para finalizar...")
	fmt.Scanln(&foo)
}
