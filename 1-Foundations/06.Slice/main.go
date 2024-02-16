package main

import "fmt"

/*
Quando lidamos com slices a forma de criacao dele é diferente do array mas a maior diferenca entre ambos é que o
slice ele nao tem tamanho definido aqui embaixo vc vera a diferenca entre a criacao de ambos:

var meuArray [3]int
slice := []int{}
*/
func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("O tamanho do slice é de %d, sua capacidade de %d e o valor: %v\n", len(slice), cap(slice), slice)
	fmt.Printf("O tamanho do slice é de %d, sua capacidade de %d e o valor: %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])
	fmt.Printf("O tamanho do slice é de %d, sua capacidade de %d e o valor: %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])
	fmt.Printf("O tamanho do slice é de %d, sua capacidade de %d e o valor: %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	//Os dois pontos quer dizer que voce esta fazendo um "corte" no seu slice.

	/*Podemos perceber que a capacidade do slice diminui quando verificamos o valor a partir da direita(2:) e para
	aumentar a capacidade do slice o proprio Go quando fazemos um append que é adicionar um novo valor ao slice,
	ele dobra a capacidade dele.*/

	slice = append(slice, 110)
	fmt.Printf("O tamanho do slice é de %d, sua capacidade de %d e o valor: %v\n", len(slice), cap(slice), slice)
}
