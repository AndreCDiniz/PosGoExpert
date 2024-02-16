package main

/*
O go é uma linguagem compilada
Para criar o arquivo binario usamos o codigo go build
e também podemos deicidir para qual sistema operacional podemos usar exemplo
goos=windows go build main.go
um site para podermos ver o sistema operacional e a arquitetura podemos usar o site da digitalocean
https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
*/
func main() {
	a := 4
	b := 2
	if a > b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		print("a")
	case 2:
		print("b")
	default:
		print("c")
	}
}
