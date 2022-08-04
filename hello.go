package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	for {

		exibeIntroducao()
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}

	}

}

func loop() {
	count := 1
	for i := 0; i < count; i++ {
		fmt.Println("loop:", i)
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	fmt.Println("Ola senhor", nome)

	versao := 1.2
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel versão é", reflect.TypeOf(versao))
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	fmt.Println(sites)

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			testaSite(site, i)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
		fmt.Println("")

	}
}

func testaSite(site string, posicao int) {
	resp, err := http.Get(site)
	//fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", posicao, "-", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", posicao, "-", site, "está com problemas", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	var sites []string
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("erro ao abrir arquivo de log:", err)
	}

	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "-online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Teve um erro:", err)
	}

	fmt.Println(string(arquivo))

}
