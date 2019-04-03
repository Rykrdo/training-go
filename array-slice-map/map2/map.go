package main

import "fmt"

func main() {
	funcESalario := map[string]float64{
		"Jose da silva": 2.1212121,
		"Pedro Ribeiro": 322.32322,
		"Amanda Nunes":  9891.9212,
		"Julita Lopes":  891329.00,
	}

	funcESalario["Barnabe"] = 53457688
	delete(funcESalario, "Inexite")
	fmt.Println(funcESalario)

	for nome, salario := range funcESalario {
		fmt.Println(nome, salario)
	}

}
