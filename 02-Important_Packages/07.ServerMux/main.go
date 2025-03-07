package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Homehandler)
	mux.Handle("/blog", blog{})
	http.ListenAndServe(":8080", mux)
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog"))
}
