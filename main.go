package main

import (
	"fmt"
)

func dadosPessoa(nome string, idade int) (int, string) {
	if idade >= 18 {
		return idade, "maior de idade"
	}
	return idade, "menor de idade"
}

func main() {
	idade, status := dadosPessoa("Bangchan", 27)
	fmt.Printf("Idade: %d - Situação: %s\n", idade, status)
}
