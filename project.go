package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	exibirDono()
	for {
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
			time.Sleep(2000 * time.Millisecond)
			os.Exit(0)

		default:
			fmt.Println("Comando inválido")
		}
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
	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.catho.com.br", "http://www.fatecbarueri.edu.br/"}

	for i, site := range sites {
		fmt.Println("Testando site", i+1)
		time.Sleep(1000 * time.Millisecond)
		verificarSite(site)
	}

}

func verificarSite(site string) {

	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregando com sucesso. status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "Apresentou algum erro. status:", resp.StatusCode)
	}
}
