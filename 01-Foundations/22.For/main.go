package main

/*Diferente de outras linguagens o go só tem o for e podemos dizer que existem 4 tipos*/
func main() {

	//for padrao
	for i := 0; i <= 10; i++ {
		println(i)
	}

	//for usando range, bastante utilziado para percorrer slices, arrays, structs on o k é o indice e o v o valor
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	//for que seria utilizado como while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//loop infinito
	for {
		println("Hello, World!")
	}
}
