package main

import "fmt"

func main() {
	funcPorletra := map[string]map[string]float64{
		"A": {
			"Luiza": 8999.90,
		},
		"B": {
			"Pedro":     90.00,
			"Adalverto": 901.3230,
			"Breno":     90.767,
		},
		"P": {
			"Jose João": 12.3208,
			"Henrique":  0.3208,
		},
	}

	for letra, nome := range funcPorletra {
		fmt.Printf("Teste %s : %s \n", letra, nome)
	}
	// 	fmt.Println(funcPorletra)
}
