package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}

	contaDoDenis.Depositar(100)

	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())
}
