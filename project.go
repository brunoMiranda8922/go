package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorar = 3
const delay = 3

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
			time.Sleep(1050 * time.Millisecond)
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
	//sites := []string{"https://random-status-code.herokuapp.com/",
	//"https://www.catho.com.br", "http://www.fatecbarueri.edu.br/"}

	sites := lerArquivo()
	for i := 0; i < monitorar; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1)
			time.Sleep(1000 * time.Millisecond)
			verificarSite(site)
			time.Sleep(delay * time.Second)
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func verificarSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro inesperado %v", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregando com sucesso. status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "Apresentou algum erro. status:", resp.StatusCode)
	}
}

func lerArquivo() []string {
	var site = []string{}
	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {

		fmt.Println("Ocorreu um erro: %v", err)

	} else {

		//fmt.Println("Arquivo", arquivo)
		sites := append(strings.Fields(string(arquivo)))

		return sites

	}

	return site
}
