package main

import "fmt"

func main() {

	//ARITMÉTICOS
	fmt.Println("-------------------")
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 1 * 2
	resto := 1 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	//ATRIBUIÇÃO
	fmt.Println("-------------------")
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println("-------------------")

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNARIOS
	fmt.Println("-------------------")
	numero := 10
	//mais 1
	numero++
	//mais 10
	numero += 10 // numero = numero + 10
	fmt.Println(numero)
	//decresmentar
	numero--
	numero -= 5
	fmt.Println(numero)
	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	//OPERADORES TERNÁRIO
	fmt.Println("-------------------")
	//não existe em go

}
