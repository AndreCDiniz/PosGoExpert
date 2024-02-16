package matematica

/*
Quando falamos em modificadores de acesso é, toda vez que a variavel, metodo, struct estiver com
a letra maiscula ela será publica ou seja, visivel a todo o projeto
caso ela seja minuscula ela só será visivel centro do mesmo pacote que no caso seria o default
*/
func Soma[T int | float64](a, b T) T {
	return a + b
}
