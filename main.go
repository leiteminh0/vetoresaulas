package main

import (
	"fmt"
)

func main() {
	var numeros [5]int
	var soma int = 0

	fmt.Println("Digite 5 números inteiros:")

	for i := 0; i < 5; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}

	for i := 0; i < 5; i++ {
		soma += numeros[i]
	}

	fmt.Println("A soma dos números é:", soma)
}
