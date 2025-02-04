package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // Aqui o channel começa vazio

	// Thread 2
	go func() {
		channel <- "Hello World!" // Aqui o channel recebe a mensagem e fica cheio
	}()

	// Thread 1
	msg := <-channel // Aqui o channel é lido e esvaziado
	fmt.Println(msg)
}
