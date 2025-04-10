package main

import (
	"fmt"
)

func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("menor que 30 anos")
	} else if age < 40 {
		fmt.Println("menor que 40 anos")
	} else {
		fmt.Println("nao é menor que 40 anos")
	}
	names := []string{"bangchan", "han", "felix", "minho", "seugmin", "changbin"}

	for index == 1 {
		fmt.Println("continue apos a posiçao", index, "e valor", value)
		continue
	}
	if index > 2 {
		fmt.Println("sair apos", index)
		break
	}
	fmt.Println("valor: ", value)
}
