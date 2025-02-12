/*Dentro do go existe os tipos porem nos tambem podemos criar nossos proprios tipos(como se fosse objetos)*/

package main

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 5
	d float32 = 1.2
	e ID      = 1
)

func main() {
	fmt.Printf("Tipo da variavel E é %T", e) //Esse %T é pra mostrar o tipo da variavel
}
