package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta := ContaCorrente{
		titular:       "Jane",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         1250.00}

	conta2 := ContaCorrente{"Joe", 222, 654321, 2540.50}

	fmt.Println(conta)
	fmt.Println(conta2)
}
