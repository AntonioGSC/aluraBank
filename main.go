package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "Erro ao sacar", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Erro ao depositar", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func main() {
	conta := ContaCorrente{
		titular:       "Jane",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         1250.00}

	conta2 := ContaCorrente{"Joe", 222, 654321, 2540.50}

	fmt.Println("=================BANCO===================")
	fmt.Println("Saldo conta 1:", conta.saldo)
	fmt.Println("Saldo conta 2:", conta2.saldo)
	fmt.Println("=========================================")
	fmt.Println(conta.Sacar(300))
	fmt.Println("=========================================")
	status, valor := conta.Depositar(300)
	fmt.Println(status)
	fmt.Println("Seu saldo atual é de:", valor)
	fmt.Println("=========================================")
	fmt.Println(conta.Transferir(1000, &conta2))
	fmt.Println("Saldo conta 1:", conta.saldo)
	fmt.Println("Saldo conta 2:", conta2.saldo)
}

// var conta3 *ContaCorrente
// conta3 = new(ContaCorrente)
// conta3.titular = "Cris"
// fmt.Println(*conta3)
