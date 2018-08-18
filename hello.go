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

}
