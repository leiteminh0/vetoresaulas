package main

import "fmt"

func main() {

	nomes := []string{"Ana", "Bruno", "Carlos", "Daniela", "Eduardo"}

	fmt.Println("Os dois primeiros:", nomes[:2])

	fmt.Println("Os dois últimos:", nomes[len(nomes)-2:])

	fmt.Println("Nome do meio:", nomes[2:3])
}
