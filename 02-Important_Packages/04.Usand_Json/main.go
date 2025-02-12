package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	//Podemos utilizar o Marshal para salvar o json em uma variavel e usar ela posteriormente
	conta := Conta{Numero: 10, Saldo: 20}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	//Outra forma de podermos fazer isso de uma forma onde ele ja receba o valor e já faça a serialização usamos o Encoder
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	//Agora iremos fazer a forma inversa, ou seja, teremos um json e queremos fazer ele voltar a struct
	//Imaginemos que os dados desse json sejam diferentes da struct que temos, por exemplo numero ficaria n e saldo s
	//para resolvermos isso podemos usar as TAGs
	jsonPuro := []byte(`{"n": 20, "s": 400}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
