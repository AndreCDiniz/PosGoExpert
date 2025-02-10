package main

import "fmt"

func main() {
	evento := []string{"evento1", "evento2", "evento3", "evento4"}
	// evento = evento[0:]
	evento = append(evento[:0], evento[1:]...)
	fmt.Println(evento)
}
