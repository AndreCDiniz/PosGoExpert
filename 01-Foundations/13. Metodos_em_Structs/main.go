package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
}

/*Para se fazer uma funcao e atrelar ela a uma struc sempre colocamos antes do nome da funcao entre parenteses uma sigla
que seria pra representar a struc e o nome da struct exemplo: (c Cliente) e depois para chamar essa funcao é so vc colocar
o atributo .metodo nesse caso, andre.Desativar()*/

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\nAtivo: %v", c.Nome, c.Ativo)
}

func main() {

	andre := Cliente{
		Nome:  "André",
		Idade: 31,
		Ativo: true,
	}

	andre.Cidade = "Recife"
	andre.Endereco.Logradouro = "Rua do Miradouro"

	fmt.Printf("O nome é: %s\nIdade: %d\nAtivo: %v\n", andre.Nome, andre.Idade, andre.Ativo)
	fmt.Printf("Endereco: %s\nCidade: %s\n", andre.Cidade, andre.Endereco.Logradouro)
	andre.Desativar()
}
