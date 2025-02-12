package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// O main é a thread principal do programa.
func main() {
	//Thread 1
	go task("A")
	//Thread 2
	go task("B")
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	//Nada aqui
	//Sair do programa
	//Esse codigo abaixo é apenas temporario para ver o resultado, gambiarra.
	time.Sleep(15 * time.Second)
}
