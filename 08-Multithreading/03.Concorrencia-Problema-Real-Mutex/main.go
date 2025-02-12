package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	/*Nessa execuçao sem o mutex podemos utilizar o go run -race onde verifica se
	temos problemas de race coditions juntamente com o apache benchmark (ab)*/
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa pagina %d vezes", number)))
	})
	http.ListenAndServe(":8888", nil)
}
