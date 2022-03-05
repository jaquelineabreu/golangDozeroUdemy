package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivos structs")

	//Podemos inserir dados assim,
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	//assim,
	usuario2 := usuario{"Mara", 20}
	fmt.Println(usuario2)

	//ou assim, caso precise de menos valores
	usuario3 := usuario{nome: "Lara"}
	fmt.Println(usuario3)

}
