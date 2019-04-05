package main

import "fmt"

func cosure() func() {
	x := 10
	var funccao = func() {
		fmt.Println("funcao", x)
	}
	return funccao
}

func main() {
	x := 20

	fmt.Println(x)

	imprimeX := cosure()
	imprimeX()
}
