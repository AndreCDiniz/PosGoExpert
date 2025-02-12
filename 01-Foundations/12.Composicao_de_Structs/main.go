package main

import "fmt"

/*Composicao de struct nada mais é do que vc usar uma struct para compor dentro de outra (fazer um encadeamento) o que as
vezes não é muito legal pois pode aumentar a complexidade do codigo.*/

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	//colocando dessa forma seria uma composicao
	Endereco
	//A forma abaixo eu estou atribuindo a variavel endereco o tipo Endereco
	//Endereco Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
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
	fmt.Printf("Endereco: %s\nCidade: %s", andre.Cidade, andre.Endereco.Logradouro)

}
