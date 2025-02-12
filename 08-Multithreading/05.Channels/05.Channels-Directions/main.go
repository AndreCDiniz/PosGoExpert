package main

import "fmt"

// Quando usamos `chan<-` estamos especificando que o canal é apenas para envio de dados.
// Isso significa que a função pode enviar dados para o canal, mas não pode recebê-los.
// Quando enviamos dados para o canal, ele pode ficar cheio se não houver uma goroutine pronta para recebê-los.
func receive(name string, hello chan<- string) {
	hello <- name
}

// Quando usamos `<-chan` estamos especificando que o canal é apenas para recebimento de dados.
// Isso significa que a função pode receber dados do canal, mas não pode enviá-los.
// Quando recebemos dados do canal, ele é esvaziado, permitindo que novos dados sejam enviados.
func read(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	hello := make(chan string)
	go receive("Hello, World!", hello)
	read(hello)
}
