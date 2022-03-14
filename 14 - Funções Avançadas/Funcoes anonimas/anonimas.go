package main

import "fmt"

//FUNÇÃO ANONIMA APENAS COM PRINT

// func main() {
// 	func() {
// 		fmt.Println("Ola Mundo!")
// 	}()
// }

//ou

//FUNÇÃO ANONIMA PASSANDO PARAMENTRO

// func main() {
// 	func(texto string) {
// 		fmt.Println(texto)
// 	}("Passando Parâmetro")
// }

//ou

//FUNÇÃO ANONIMA COM PARAMENTRO E RETURN

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)
}
