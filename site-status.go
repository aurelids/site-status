package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 2
const delay = 3

func main() {
	
	exibeIntroducao()
	registraLog("site-falso", false)
	
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	var nome string = "Douglas"
	var idade int = 24
	versao := 1.1 // nao precisa declarar e nem dizer var se colocar :=
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa esta na versao", versao)
}

func leComando() int {
	reader := bufio.NewReader(os.Stdin)
	comandoLido, _ := reader.ReadString('\n')
	comandoLido = strings.TrimSpace(comandoLido)
	comando, err := strconv.Atoi(comandoLido)
	if err != nil {
		fmt.Println("Erro ao ler o comando:", err)
		return -1
	}
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br", "https://www.google.com", "https://instagram.com", "https://ge.globo.com" }

	sites := leSitesDoArquivo()
	
	fmt.Println(sites)

	// for i:= 0 ; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	
	for i:= 0; i < monitoramentos ; i++ {

		for i, site := range sites {
				fmt.Println("Testando site", i, ":", site)
				testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}
	fmt.Print("")
}

func testaSite(site string){
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo () []string {

	var sites []string
	 arquivo, err := os.Open("sites.txt")
	// arquivo, err := os.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')						// \n isso quer dizer que é para ele ler o documento até a quebra de linha (ou seja, ler apenas uma linha)
		linha = strings.TrimSpace(linha)							// removendo o '\n' que ele estava adicionando
		
		sites = append(sites, linha)								// colocando as linhas dentro da string
		if err == io.EOF {											// EOF = end of file, significa que acabou o arquivo
			break													// sai do for
		}
	}

	arquivo.Close()													//boa pratica em fechar o arquivo.
	return sites
}

func registraLog(site string, status bool){
	
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)					// funcao propria (0666)

	if err != nil {
		fmt.Println(err)
	}
	
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")			// a data tem que consultar na documentação

	arquivo.Close()
}

func imprimeLogs(){
	
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

