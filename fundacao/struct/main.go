package main

import "fmt"

type Conta struct {
	saldo int
}

func (c *Conta) Depositar(deposito int) {
	c.saldo += deposito
}

func main() {
	var conta = Conta{100}
	conta.Depositar(200)
	fmt.Println(conta.saldo)
}