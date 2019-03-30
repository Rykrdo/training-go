package main

import "fmt"

func main() {

	numeros := [...]int{1, 3, 7, 8, 10, 11}

	for i, numero := range numeros {
		fmt.Printf("O número %d = %d \n", i, numero)
	}

}
