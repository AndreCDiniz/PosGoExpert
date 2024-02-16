package main

import (
	"21.Pacotes_e_Modulos/matematica"
	"fmt"
	"github.com/google/uuid"
)

/*
Normalmente a primeira coisa que fazemos quanto vamos programar em go é fazer o pacote com o gomod
por boas praticas ao criar o modulo usando o go mod init é interessante colocarmos o diretorio do github
para que quando alguma pessoa quiser usar o modulo é so usar o endereço do git
*/
func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("O resultado:", s)

	fmt.Println(uuid.New())
}
