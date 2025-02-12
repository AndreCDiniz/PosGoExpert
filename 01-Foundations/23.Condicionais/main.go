package main

// dentro do go nos possuimos o if e switch e o select(esse vai ser dado mais a frente)
func main() {

	a := 3
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
