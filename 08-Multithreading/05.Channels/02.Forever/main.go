package main

// Thread 1
func main() {

	forever := make(chan bool) // Aqui o channel começa vazio

	// Thread 2
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		// Para resolver o deadlock, podemos fechar o channel ou enviar um valor
		// close(forever) // Fechar o channel indica que não haverá mais envios de dados, evitando deadlock
		forever <- true // Aqui o channel recebe a mensagem e fica cheio
	}()

	/*Aqui o channel é lido e esvaziado e se ele nao for lido,
	  o programa fica em execução para sempre dando um deadlock */
	<-forever
}
