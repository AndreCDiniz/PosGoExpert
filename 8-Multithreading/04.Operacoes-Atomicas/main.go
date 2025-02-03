package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	/*Usar mutex o tempo inteiro funciona mas o Go ele tem recursos melhores do que isso
	um deles é soma atomica, que é fazer incremento de alguma coisa de forma isolado sem
	interferencia de Goroutines mesmo quando varias delas estão tentando acessar o mesmo
	valor ao mesmo tempo. */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// internamente ele já esta usando o mutex de lock e unlock
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa pagina %d vezes", number)))
	})
	http.ListenAndServe(":8888", nil)
}
