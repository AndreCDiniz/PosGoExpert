package main

/*
Para um bom entendimento, quando pensarmos nos valores das variaves na funcao soma o a que está sendo passado pela
variavel minhaVar1 na verdade está mandando uma cópia e caso eu precise enviar o enderecamento de memoria eu faco de
uma forma diferente.
*/
func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

/*apos adicionar o * agora estamos querendo o endereço de memoria ou seja, estamos transformando isso em um ponteiro*/

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	//agora devemos colocar o endereço de memoria das variaveis com o & .
	soma(&minhaVar1, &minhaVar2)
	println(minhaVar1) //antes essa variavel saia 10 e agora ela vai sair 50 porque agora nos mudamos para endereco de memoria
}

/*Quando não usar ponteiro?

- Quando quer passar apenas copia de dados ou utilizacao deles. Exemplo metodo de soma acima

Quando usar?

- Quando quiser tornar por alguma motivos os valores para mutaveis  */
