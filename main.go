package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Erro ao sacar"
	}
}

func main() {
	conta := ContaCorrente{
		titular:       "Jane",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         1250.00}

	fmt.Println(conta.saldo)
	fmt.Println(conta.Sacar(300))
	fmt.Println(conta.saldo)

	// conta2 := ContaCorrente{"Joe", 222, 654321, 2540.50}
	// fmt.Println(conta2)

	// var conta3 *ContaCorrente
	// conta3 = new(ContaCorrente)
	// conta3.titular = "Cris"
	// fmt.Println(*conta3)
}
