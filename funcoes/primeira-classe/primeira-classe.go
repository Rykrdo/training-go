package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	test := soma(2, 3)

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(test)
	fmt.Println(sub(10, 5))
}
