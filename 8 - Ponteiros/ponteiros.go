package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var varivale2 int = variavel1 //Passado valor por cópia

	fmt.Println(variavel1, varivale2)

	varivale2++
	fmt.Println(variavel1, varivale2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int // esse guarda um valor inteiro
	var ponteiro *int //esse guarda o valor de um endereço de memoria

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)  // mostra o valor e o endereço da memoria
	fmt.Println(variavel3, *ponteiro) // mostra o valor e a desreferenciação ou desfazendo a referencia

}
