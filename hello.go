package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	mensagem := "Hello mundo"
	fmt.Println(mensagem)
	fmt.Println("Mensagem é do tipo:", reflect.TypeOf(mensagem))
	fmt.Println("Digite um número entre 1 á 10")

	//caminho, posição na mémoria que a variavel esta

	var numero int
	fmt.Scan(&numero)
	fmt.Println("Posição na memória é:", &numero)

	//array

	var array [4]int
	array[0] = 22
	array[1] = numero
	array[2] = 0
	array[3] = array[0]
	fmt.Println("Retorno do array", array)
	fmt.Println("Tamanho do array:", len(array))
	fmt.Println("Capacidade do array é:", cap(array))
	//Slice -> abstração de um array

	slice := []string{"Bruno", "Marcio", "Barreto"}
	fmt.Println("Retorno do slice:", slice, "Tamanho do slice é de ", len(slice))
	fmt.Println("Capacidade do slice é:", cap(slice))
	fmt.Println("Adicionando mais um elemento no slice..")
	time.Sleep(3000 * time.Millisecond)
	slice = append(slice, "Mario")
	fmt.Println("Retorno do slice:", slice, "Tamanho do slice é de ", len(slice))
	fmt.Println("Capacidade do slice é:", cap(slice))
}
