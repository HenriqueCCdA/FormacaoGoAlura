package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)

	return comandoLido
}
