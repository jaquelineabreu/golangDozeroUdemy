package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	//IF
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//IF INIT
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Print("Numero é menor que zero")
	}

}
