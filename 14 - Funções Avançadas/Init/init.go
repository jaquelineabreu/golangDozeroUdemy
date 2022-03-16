package main

import "fmt"

var n string

func init() {
	fmt.Println("Executando a função init")
	n = "Inicializando com init"
}

func main() {
	fmt.Println("Função mais sendo executada")
	fmt.Println(n)

}
