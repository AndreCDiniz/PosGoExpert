package main

/*Logo de cara, quando nao existia Generis precisariamos repetir a funcao abaixo duas vezes caso eu queria
fazer a soma tanto para float como para int, coisa feia...*/
/*func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}*/

/*
Mas depois das novas versões 1.18 o Go agora aceita genericos e logo abaixo um exemplo de como seria os dois
metodos acima
*/
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

/*
Para deixar a parte do codigo mais elegante, podemos utilizar constraints que serviria para
substituir o int | float no exemplo da funcao.

Para fazer o MyNumber do tipo int ser aceito aqui devemos adicionar o ~ na frente da variavel
*/
type Number interface {
	~int | ~float64
}

type MyNumber int

/*
Quando vamos fazer comparaçoes de variaveis utilizando generics não devemos usar o any, ou a variavel
myNumber e sim o comparable que ele só ira fazer a comparação se as variaves forem do mesmo tipo. MAS TEM UM POREM:
o comparable ele compara a igualdade das variaveis, se caso voce for usar um > ele nao ira funcionar
Mas também existem um pacote no Go chamado de constraints link: https://pkg.go.dev/golang.org/x/exp/constraints
*/
func Comparar[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"André": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"André": 1000.30, "João": 2000.20, "Maria": 300.0}

	println(Soma(m))
	println(Soma(m2))

	/*E se por um acaso eu quiser usar um valor MyNumber do tipo int? A primeiro momento vai dar problema
	mas existe uma forma do go aceitar isso, olhe mais acima.*/
	m3 := map[string]MyNumber{"André": 1000, "João": 2000, "Maria": 8000}
	println(Soma(m3))

	print(Comparar(10, 10))

}
