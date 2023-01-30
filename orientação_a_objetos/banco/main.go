package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDaGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilvia.Transferir(200, &contaDaGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDaGustavo)

}
