package main

import (
	"fmt"
	"sync"
)

// O WaitGroup é usado para esperar que todas as goroutines terminem sua execução.
// Isso é útil quando queremos garantir que todas as operações assíncronas sejam concluídas antes de continuar.
func main() {
	ch := make(chan int) // Cria um canal para enviar inteiros.
	wg := sync.WaitGroup{}
	wg.Add(10)         // Define o número de operações que o WaitGroup deve esperar.
	go publish(ch)     // Inicia a goroutine publish para enviar dados ao canal.
	go reader(ch, &wg) // Inicia a goroutine reader para ler dados do canal. Usamos 'go' aqui porque estamos usando WaitGroup para sincronização.
	wg.Wait()          // Espera até que todas as operações do WaitGroup sejam concluídas.
}

// A função reader lê dados do canal até que ele seja fechado.
func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch { // O range lê dados do canal até que ele seja fechado.
		fmt.Printf("Received %d\n", x) // Imprime o valor recebido do canal.
		wg.Done()                      // Indica que uma operação foi concluída.
	}
}

// A função publish envia dados ao canal.
func publish(ch chan int) {
	for i := 0; i < 10; i++ { // Envia os números de 0 a 9 ao canal.
		ch <- i
	}
	close(ch) // Fecha o canal para indicar que não haverá mais envios. Se não fecharmos o canal, a função reader entrará em deadlock.
}
