/*Um array é uma variavel que tem tamanho limitado */

package main

import "fmt"

func main() {
	var meuArray [3]int //a criaçao desse array é de 3 posicoes e nao se pode mais alterar isso sempre comeca em 0

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 20

	/*Nota importante, o pelo array sempre comecar na posicao 0 sempre para ver a ultima posicao dele é o tamanho dele -1
	ou seja, se eu tenho um array de 10 posicoes a ultima posicao dele será 9*/

	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])

	/*Para percorrer um array nos utilizamos o for*/

	for i, v := range meuArray {
		fmt.Printf("A posição do indice é %d e o valor dele é %d:\n", i, v) //D nesse caso quer dizer digito.
	}

}
