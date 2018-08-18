package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibirDono()
	exibirMenu()
	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando..")
		iniciarMonitoramento()
	case 2:
		fmt.Println("Logs..")
	case 0:
		fmt.Println("Saindo..")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido")
	}
}

func exibirDono() {
	nome := "Bruno"
	fmt.Println("Você está logado com o usuário", nome)
}

func exibirMenu() {
	fmt.Println("1- Para monitorar")
	fmt.Println("2- Para ver logs")
	fmt.Println("0- Para sair")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Você selecionou o comando:", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	site := "http://www.catho.com"
	resp, _ := http.Get(site)
	fmt.Println(resp)

}
