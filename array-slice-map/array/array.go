package main

import "fmt"

func main() {
	// estrutura homogenea
	var notas [4]float64

	// fmt.Println(notas)

	notas[0], notas[1], notas[2], notas[3] = 5.9, 10.00, 7.9, 6.0

	// fmt.Println(notas[0])
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]

	}
	media := total / float64(len(notas))
	fmt.Printf("A media é %.2f", media)
}
