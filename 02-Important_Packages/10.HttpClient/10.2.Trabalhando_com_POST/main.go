package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}
	jsonVar := []byte(`{"name":"Andre","age":31}`)
	resp, err := c.Post("https://www.google.com", "application/json", bytes.NewBuffer(jsonVar))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
