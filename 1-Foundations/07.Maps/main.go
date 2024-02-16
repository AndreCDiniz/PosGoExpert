/*Dentro do go também podemos utilizar hashtables mas sao mais conhecidas como map onde temos uma chave e um valor*/

package main

import "fmt"

func main() {

	/*Temos um exemplo abaixo de como criar um map com seus devidos valores*/
	salarios := map[string]int{"André": 5000, "Joao": 7000, "Shara": 10000}
	fmt.Println(salarios)

	//Imprimir o valor de uma chave
	fmt.Println(salarios["Joao"])

	//Deletar uma chave do mapa
	delete(salarios, "Joao")

	//adicionar uma nova chave ao mapa
	salarios["Jo"] = 20000

	fmt.Println(salarios)

	//Existe tambem outras formas de instanciar um map
	salarioTest1 := make(map[string]int)
	salariosTest2 := map[string]int{}

	fmt.Println(salarioTest1, salariosTest2)

	//Se caso eu queira percorrer o map, eu pode fazer de outra forma
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é de %d\n", nome, salario)
	}

	//Caso voce nao queira imprimir o nome por exemplo voce pode usar o blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é de: %d\n", salario)
	}

	for nome, _ := range salarios {
		fmt.Printf("O nome é: %s\n", nome)
	}
}
