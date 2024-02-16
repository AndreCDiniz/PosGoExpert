package main

import "fmt"

//Ponteiros é o enderecamento da memoria

func main() {

	//Memoria -> Endereco -> Valor
	a := 10
	fmt.Println(a)
	//para poder pegar o endereco da memoria de uma variavel usamos o &
	fmt.Println(&a)

	/*Imaginemos que precisamos alterar o valor de a porem utilizando o endereco de memoria e não o atributo a, como
	fazemos isso?*/
	//Toda vez que colocamos o * é porque estamos indo para o endereco de memoria.
	var ponteiro *int = &a
	fmt.Println(ponteiro)
	//nesse caso pegamos o enderecamento do ponteiro que antes era 10 e agora estamos colocando o valor de 20
	*ponteiro = 20
	fmt.Println(a)

	//nessa caso criamos uma variavel e apontamos ela para o endereco de a ou seja, b seria um ponteiro
	b := &a
	fmt.Println(b)

	//agora como faco para ver o valor atraves do b? É só utilizar o *
	fmt.Println(*b)

	//aqui eu estou mudando o valor que antes era 20 para 30
	*b = 30

	//se formos olhar o valor de a agora ele tambem estara 30
	println(a)
}
