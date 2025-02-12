package main

type Conta struct {
	saldo int
}

/*
A partir do momento em que eu coloco o ponteiro em Conta, eu estou dizendo que o valor que ele ira receber nao sera
mais uma cópia e sim o valor que está no endereço de memoria
*/
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	//se formos rodar a primeira vez sem usar o ponteiro na funcao simular (*) esse valor sempre sera 100
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
