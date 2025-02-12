package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ViaCEP represents the data structure for the response from the ViaCEP API.
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		// Faz uma solicitação GET ao serviço web viacep.com.br para obter informações sobre o CEP
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		// Caso algo dê errado com a solicitação HTTP, o erro é impresso no stderr
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		// Garante que o corpo da resposta seja fechado no final da execução do programa
		defer req.Body.Close()

		// Lê todos os dados do corpo da resposta
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		// Inicializa a estrutura de dados ViaCEP para armazenar a resposta
		var data ViaCEP

		// Tenta desestruturar/parsear o corpo da resposta JSON para a estrutura ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}

		// Tenta criar um novo arquivo chamado cidade.txt
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		// Garante que o arquivo seja fechado no final da execução do programa
		defer file.Close()

		// Escreve as informações de CEP, localidade e UF no arquivo cidade.txt
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade: ", data.Localidade)
	}
}
