package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	conta := contas.ContaCorrente{
		Titular:       "Jane",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         1250.00}

	conta2 := contas.ContaCorrente{"Joe", 222, 654321, 2540.50}

	fmt.Println("=================BANCO===================")
	fmt.Println("Saldo conta 1:", conta.Saldo)
	fmt.Println("Saldo conta 2:", conta2.Saldo)
	fmt.Println("=========================================")
	fmt.Println(conta.Sacar(300))
	fmt.Println("=========================================")
	status, valor := conta.Depositar(300)
	fmt.Println(status)
	fmt.Println("Seu Saldo atual é de:", valor)
	fmt.Println("=========================================")
	fmt.Println(conta.Transferir(1000, &conta2))
	fmt.Println("Saldo conta 1:", conta.Saldo)
	fmt.Println("Saldo conta 2:", conta2.Saldo)
}

// var conta3 *ContaCorrente
// conta3 = new(ContaCorrente)
// conta3.titular = "Cris"
// fmt.Println(*conta3)
