package main

import (
	"fmt"
	"time"
)

// Simula um trabalhador que processa dados recebidos pelo canal.
// Cada trabalhador imprime o valor recebido e simula um tempo de processamento com Sleep.
func worker(workerID int, data chan int) {
	for x := range data { // O range lê dados do canal até que ele seja fechado.
		fmt.Printf("Worker %d got %d\n", workerID, x) // Imprime o ID do trabalhador e o valor recebido.
		time.Sleep(time.Second)                       // Simula um tempo de processamento de 1 segundo.
	}
}

// A função main é a thread principal que cria o canal e inicia as goroutines dos trabalhadores.
func main() {
	data := make(chan int) // Cria um canal para enviar inteiros.
	WorkersCount := 10000  // Define o número de trabalhadores.

	// Inicia as goroutines dos trabalhadores.
	for i := 0; i < WorkersCount; i++ {
		go worker(i, data) // Inicia uma goroutine worker para cada trabalhador.
	}

	// Envia dados para os trabalhadores processarem.
	for i := 0; i < 100000; i++ {
		data <- i // Envia o valor i para o canal.
	}
}
