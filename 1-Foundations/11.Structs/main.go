package main

import "fmt"

/*strucs são a forma que mais se assemelha a uma classe na orientacao a objetos porem a linguagem go nao é orientada a
objetos. Basicamente uma struct é uma estrutura de dados onde dentro dela pode ser armazenado varios tipos diferentes.*/

type Cliente struct {
	nome  string
	idade int
	ativo bool
}

func main() {

	andre := Cliente{
		nome:  "André",
		idade: 31,
		ativo: true,
	}

	fmt.Printf("O nome é: %s\nIdade: %d\nAtivo: %v\n", andre.nome, andre.idade, andre.ativo)

	//Também é possivel mudar dados de uma struct por exemplo

	andre.ativo = false

	fmt.Println(andre.ativo)
}
