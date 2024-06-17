package main

import (
	"io"
	"net/http"
)

// Essa seria a forma mais simples de voce fazer chamadas HTTP no go.
func main() {

	// **1. Requisição GET:**

	// Enviamos uma requisição GET para a página inicial do Google.
	// A função `http.Get` retorna um objeto `*http.Response` e um valor `error`.
	// Se a requisição for bem-sucedida, o objeto `*http.Response` conterá informações
	// sobre a resposta, como o código de status e o corpo da resposta.
	// Se a requisição falhar, o valor `error` conterá a causa do erro.
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	//Após a requisição feita a acima nos precisamos pegar esse body e jogar pra dentro de uma response
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	//Importante usar o string aqui pois se não for utilizado iremos receber um hash bem bizarro
	println(string(res))

	//E depois disso fechamos a conexao
	req.Body.Close()
}
