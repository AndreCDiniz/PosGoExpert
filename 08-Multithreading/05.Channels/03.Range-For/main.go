package main

import "fmt"

// A função main é a thread principal que cria o canal e inicia as goroutines.
func main() {
	ch := make(chan int) // Cria um canal para enviar inteiros.
	go publish(ch)       // Inicia a goroutine publish para enviar dados ao canal.
	reader(ch)           // Chama a função reader para ler dados do canal. Não usamos 'go' aqui para garantir que a função main espere até que todos os valores sejam lidos antes de terminar.
}

// A função reader lê dados do canal até que ele seja fechado.
func reader(ch chan int) {
	for x := range ch { // O range lê dados do canal até que ele seja fechado.
		fmt.Printf("Received %d\n", x) // Imprime o valor recebido do canal.
	}
}

// A função publish envia dados ao canal.
func publish(ch chan int) {
	for i := 0; i < 10; i++ { // Envia os números de 0 a 9 ao canal.
		ch <- i
	}
	close(ch) // Fecha o canal para indicar que não haverá mais envios. Se não fecharmos o canal, a função reader entrará em deadlock.
}
