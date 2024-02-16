/*Existe uma forma de usar shorthand and inves de usar o tipo da variavel podemos usar a inferencia usando :=*/

package main

var (
	b bool
	c int
	d float32
)

func main() {

	a := "X" //ele ja vai entender que é uma String uma vez usado a inferencia := nao se usa pra mesma variavel novamente
	a = "y"

	println("Hello World " + a)
	println(b)
	println(c)
	println(d)
}
