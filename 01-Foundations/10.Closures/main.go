package main

import (
	"fmt"
)

/*As funcoes closures são quando voce usa uma funcao dentro de uma funcao isso é bastante usado em routines no go*/
func main() {
	total := func() int {
		return sum(5, 50, 911, 213131, 312, 312, 31, 5, 135, 32, 53, 5, 634, 7, 5) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
