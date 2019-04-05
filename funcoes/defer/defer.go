package main

import "fmt"

func obteValorAprovado(numero int) int {
	defer fmt.Println("Fim")

	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor muito abaixo")
	return numero
}

func main() {
	fmt.Println(obteValorAprovado(23))
}
