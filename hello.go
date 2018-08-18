package main

import (
	"fmt"
	"reflect"
)

func main() {
	mensagem := "Hello mundo"
	fmt.Println(mensagem)
	fmt.Println("Mensagem é do tipo:", reflect.TypeOf(mensagem))
	fmt.Println("Digite um número entre 1 á 10")

	var numero int
	fmt.Scan(&numero)
	fmt.Println(&numero)

	var array [4]int
	array[0] = 22
	array[1] = numero
	array[2] = 0
	array[3] = array[0]

	fmt.Println("Retorno do array", array)
}
