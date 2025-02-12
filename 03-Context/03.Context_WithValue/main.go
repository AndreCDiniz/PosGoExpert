package main

import (
	"context"
	"fmt"
)

// O contexto pode passar dados de um lado para outro.
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

// Por convensão o contexto sempre vem como primeiro parametro na função.
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
