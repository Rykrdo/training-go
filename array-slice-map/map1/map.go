package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[33534] = "Bruno"
	aprovados[54354] = "Angelico"
	aprovados[21212] = "Renan"
	aprovados[67953] = "Ana"
	aprovados[24332] = "Roberta"

	// fmt.Println(aprovados)

	fmt.Println("Usuários cadastrado:\n")
	for cpf, nome := range aprovados {
		fmt.Println(cpf, nome)
	}
	fmt.Println(aprovados[54354])
	delete(aprovados, 54354)
	fmt.Println(aprovados[54354])

}
