/*As funcoes dentro do Go funcionam da mesma forma que todas as linguagens.*/

package main

import (
	"errors"
	"fmt"
)

// Nesse caso a funcao main não tem nenhum retorno
func main() {
	fmt.Println(sum(5, 50))
	fmt.Println(sum(5, 10))

	fmt.Println(sub(5, 6))
	fmt.Println(sub(10, 5))

	/*outra forma que podemos usar esse err é quando criamos a variavel
	nesse caso por ja existir o metodo sub eu posso utilizar o valor inteiro
	que é o primeiro parametro e o segundo um erro e fazer uma validacao desse error com if*/

	valor, err := sub(5, 10)
	if err != nil {
		fmt.Println("\nComo o numero é abaixo de 0 ele vai dar o error")
		fmt.Println(err)
	}
	fmt.Println(valor)
}

/*
Nessa funcao recebemos dois parametros do mesmo tipo, por isso,
ao inves de voce fazer a int e b int voce pode coloar de uma forma só
*/

/*As funcoes no go também podem retornar mais de um valor, como no exemplo abaixo podemos usar um bool*/
func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

/*
Como dentro do go nao existe try catch uma forma de utilizarmos dois retorno é pra tratar de erros usando err
o nil seria vazio, nulo ou nao existe
*/
func sub(a, b int) (int, error) {
	if a-b < 0 {
		return a - b, errors.New("O numero é negativo!")
	}
	return a - b, nil
}
