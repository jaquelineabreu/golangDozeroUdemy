package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//MAP NORMAL
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	//MAP ANINHADO
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": " João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1 ",
		},
	}

	fmt.Println(usuario2)
	//DELETAR MAP
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//ADICIONAR MAP
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)

}
