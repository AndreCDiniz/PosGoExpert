package main

import (
	"fmt"
)

// Nesse caso a funcao main não tem nenhum retorno
func main() {
	fmt.Println(sum(5, 50, 911, 213131, 312, 312, 31, 5, 135, 32, 53, 5, 634, 7, 5))
}

/*
As funcoes variadicas no go podem ser usadas quando nao sabemos a quantidade de parametros que iremos usar
em uma funcao desde que ela seja do mesmo tipo. No exemplo abaixo nao sabemos quantos numeros podem ser passados
para a soma
*/
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
