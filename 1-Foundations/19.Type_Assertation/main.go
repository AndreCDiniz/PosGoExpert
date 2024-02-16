package main

import "fmt"

/*
Por vezes quando utilizamos interfaces vazias a gente pode forçar que esse tipo "vazio" ser transformado
para que o Go saiba o que sera utilizado
*/
func main() {

	var minhavar interface{} = "André Diniz"

	/*Quando queremos transformar a nossa variavel em outra usamos esse .(tipo)*/
	println(minhavar.(string))

	/*Nesse caso abaixo quando iremos imprimir o res vai ser uma string e o ok um boolean
	onde se o valor da minha var for string ele vai retornar o nome da variavel (res) e o resultado de ok*/
	res, ok := minhavar.(string)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	/*Nesse caso, eu estou tentando fazer o assertation de uma string para int e não é possivel entao ele vai
	retornar 0 no res e o ok será false*/
	res2, ok := minhavar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res2, ok)

	/*E se voce for louco de nao usar o ok logo apos o res o isso vai dar um problema e causar um PANIC*/
	res3 := minhavar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res3, ok)
}
