package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// A função main é a thread principal que cria os canais e inicia as goroutines.
func main() {

	// Estrutura para armazenar mensagens com um ID e um texto.
	type Message struct {
		ID  int64
		Msg string
	}

	c1 := make(chan Message) // Cria um canal para mensagens do RabbitMQ.
	c2 := make(chan Message) // Cria um canal para mensagens do Kafka.
	var i int64 = 0          // Variável para gerar IDs únicos para as mensagens.

	// Simula recebimento de mensagens do RabbitMQ.
	go func() {
		for {
			// Incrementa o ID de forma atômica para evitar problemas de concorrência.
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second) // Simula um atraso no recebimento da mensagem.
			msg := Message{ID: i, Msg: "Hello from RabbitMQ"}
			c1 <- msg // Envia a mensagem para o canal c1.
		}
	}()

	// Simula recebimento de mensagens do Kafka.
	go func() {
		for {
			// Incrementa o ID de forma atômica para evitar problemas de concorrência.
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second) // Simula um atraso no recebimento da mensagem.
			msg := Message{ID: i, Msg: "Hello from Kafka"}
			c2 <- msg // Envia a mensagem para o canal c2.
		}
	}()

	// Loop principal que processa as mensagens recebidas dos canais.
	for {
		select {
		case msg := <-c1: // Caso receba uma mensagem do canal c1 (RabbitMQ).
			fmt.Printf("Received from RabbitMQ %d: %s\n", msg.ID, msg.Msg)
		case msg := <-c2: // Caso receba uma mensagem do canal c2 (Kafka).
			fmt.Printf("Received from Kafka %d: %s\n", msg.ID, msg.Msg)
		case <-time.After(time.Second * 2): // Timeout se nenhuma mensagem for recebida em 2 segundos.
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
