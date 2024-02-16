package main

import "fmt"

/*
A ideia da interface é que quando criamos e colocamos um método dentro dela, significa que todos os lugares
que utilizarem o metodo dentro dela estará implementando a interface. Porem, se nao colocar nenhum metodo dentro dela
ou seja, vazia, quer dizer que essa interface implementa tudo.
*/
func main() {

	/*Contudo, é bom utilizar isso com muito cuidado para sempre usarmos a tipagem do Go.
	E hoje existe o recurso de Generics para evitar utilizar isso.*/
	var x interface{} = 10
	var y interface{} = "Hello World"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T, e o valor é %v\n", t, t)
}
