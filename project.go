package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
			imprimirLogs()
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
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregando com sucesso. status:", resp.StatusCode)
		logsArquivo(site, true)
	} else {
		fmt.Println("Site:", site, "Apresentou algum erro. status:", resp.StatusCode)
		logsArquivo(site, false)
	}
}

func lerArquivo() []string {
	var site = []string{}

	//arquivo, err := os.Open("sites.txt")
	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {

		fmt.Println("Ocorreu um erro: %v", err)

	} else {

		//fmt.Println("Arquivo", arquivo)
		sites := strings.Fields(string(arquivo))
		return sites

	}

	return site
}

//Outra forma de fazer uma função para abrir e ler um arquivo.

// func lerArquivo() []string {

// 	var sites []string

// 	arquivo, err := os.Open("sites.txt")
// 	if err != nil {
// 		fmt.Println("Ocorreu um erro:", err)
// 	}

// 	leitor := bufio.NewReader(arquivo)

// 	for {
// 		linha, err := leitor.ReadString('\n')
// 		linha = strings.TrimSpace(linha)
// 		sites = append(sites, linha)
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	arquivo.Close()

// 	return sites
// }

func logsArquivo(site string, status bool) {
	logs, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}

	//data := time.Now().Format("02/01/2006")
	data := time.Now().Format(time.RFC850)
	fmt.Println(data)

	logs.WriteString("Data: " + data + "-" + site + "- online: " + strconv.FormatBool(status) + "\n")

	logs.Close()
}

func imprimirLogs() {
	mostrar, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(mostrar))

}
