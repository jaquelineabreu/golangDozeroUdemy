package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//ARRYS
	var array1 [5]string // declaração
	fmt.Println(array1)
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicação 1", "Posicação 2", "Posicação 3", "Posicação 4", "Posicação 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//VERIFICANDO TIPOS
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	//Pegando fatias
	fmt.Println(array2)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	//ARRAYS INTERNO---------------------------------
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //para saber o tamanho
	fmt.Println(cap(slice3)) //para saber a capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //para saber o tamanho
	fmt.Println(cap(slice3)) //para saber a capacidade

	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //para saber o tamanho
	fmt.Println(cap(slice3)) //para saber a capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
