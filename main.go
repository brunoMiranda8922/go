package main

import "fmt"

func main() {
	nome, idade := nomeEIdade()
	fmt.Println("Hello eu me chamo", nome, "e tenho", idade, "anos")

	//fmt.Println("O tipo é", reflect.TypeOf(nome)) Para ver o tipo da String

}

func nomeEIdade() (string, int) {
	nome := "Bruno Miranda"
	idade := 22

	return nome, idade
}
