package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		fmt.Printf("Done() chamado pela task %s, iteração %d\n", name, i)
		wg.Done()
	}
}

// O main é a thread principal do programa.
func main() {
	//Waitgroup serve para esperar todas as threads terminarem nesse caso seria 25 threads
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	//Thread 1
	go task("A", &waitGroup)
	//Thread 2
	go task("B", &waitGroup)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			fmt.Printf("Done() chamado pela anonymous, iteração %d\n", i)
			waitGroup.Done()
		}
	}()
	fmt.Println("Aguardando 5 Done()...")
	waitGroup.Wait()
	fmt.Println("Recebeu 5 Done(), programa terminando...")
}
