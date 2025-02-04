package main

// Pode haver momentos em que eu queira que o canal receba mais do que uma informaçao ao mesmo tempo. Ou seja ao inves de travar o canal, ele pode receber mais de uma informaçao.
// E isso pode ser solucinado com o buffer. O buffer é um segundo parametro que eu posso passar para o make, que é o tamanho do buffer.
// Porem não é recomendado usar buffer, pois irá aloca memória e pode causar problemas de concorrência.
func main() {

	ch := make(chan string, 2) // Cria um canal com buffer de tamanho 2.

	ch <- "Hello" // Envia a string "Hello" para o canal.
	ch <- "World" // Envia a string "World" para o canal.

	println(<-ch) // Recebe a string "Hello" do canal e imprime.
	println(<-ch) // Recebe a string "World" do canal e imprime.

}
