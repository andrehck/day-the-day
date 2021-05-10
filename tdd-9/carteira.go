package carteira

import "fmt"

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string { //altero no pacote fmt o time string para receber o valor do return
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Retirar(quantidade Bitcoin) {
	c.saldo -= quantidade
}
