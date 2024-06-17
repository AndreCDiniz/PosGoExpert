package main

import "fmt"

// O defer é sempre executado após a execução do bloco em que ele foi declarado
func main() {
	defer fmt.Println("Primeira execução")
	fmt.Println("Segunda execução")
	fmt.Println("Terceira execução")
}
