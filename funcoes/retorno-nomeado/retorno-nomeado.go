package main

import "fmt"

func trocar(p1, p2 int) (segungo int, primeiro int) {
	segungo = p2
	primeiro = p1
	return
}
func main() {
	x, y := trocar(1, 2)
	fmt.Println(x, y)

	primeiro, segundo := trocar(10, 56)

	fmt.Println(primeiro, segundo)
}
