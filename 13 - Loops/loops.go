package main

import (
	"fmt"
	"time"
)

func main() {

	//FORMA 1

	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)

	// }

	// fmt.Println(i)

	//FORMA 2

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)

	// }

	//FORMA 3
	//Iterrar array ou slice
	nomes := []string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//ou só apontando o indice
	for nome := range nomes {
		fmt.Println(nome)
	}

	//ou ocultando o indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//iterando string
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	//Iterando Map
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
