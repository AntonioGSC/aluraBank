package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	conta := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Jane",
			CPF:       "123.123.123-12",
			Profissao: "Dev Front-End"},
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	cliente2 := clientes.Titular{
		Nome:      "Joe",
		CPF:       "789.789.789-89",
		Profissao: "Dev Back-End",
	}

	conta2 := contas.ContaCorrente{
		Titular:       cliente2,
		NumeroAgencia: 222,
		NumeroConta:   654321,
	}

	// conta2 := contas.ContaCorrente{clientes.Titular{"Joe", "789.789.789-89", "Dev Back-End"}, 222, 654321, 2540.50}

	fmt.Println("=================BANCO===================")
	conta.Depositar(1500.25)
	conta2.Depositar(2250.99)
	fmt.Println("saldo conta 1:", conta.ObterSaldo())
	fmt.Println("saldo conta 2:", conta2.ObterSaldo())
	fmt.Println("=========================================")
	fmt.Println(conta.Sacar(300))
	fmt.Println("=========================================")
	status, valor := conta.Depositar(300)
	fmt.Println(status)
	fmt.Println("Seu saldo atual é de:", valor)
	fmt.Println("=========================================")
	fmt.Println(conta.Transferir(1000, &conta2))
	fmt.Println("saldo conta 1:", conta.ObterSaldo())
	fmt.Println("saldo conta 2:", conta2.ObterSaldo())
}

// var conta3 *ContaCorrente
// conta3 = new(ContaCorrente)
// conta3.titular = "Cris"
// fmt.Println(*conta3)
