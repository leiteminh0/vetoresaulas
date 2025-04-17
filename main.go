package main

import (
"fmt"
)

func dividir (dividendo int, divisor int) (int, string){
	if dividir == 0 {
		return 0, "Erro na divisão por zero"
}
return dividendo / divisor, "Sem erro"
}

func main() {
resultado, erro := dividir(10,2)

if erro != "Sem erro" {
	fmt.Printf(erro)
} else {
	fmt.Println("O resultado da divisão é:", resultado, erro)
}

func operaçãoBasica(a int, b int) (int, int, int) {
	soma := a + b
	multiplicacao := a * b
	subtracao := a - b 
	return soma, multiplicacao, subtracao
}
soma, subtracao, multiplicacao := operaçãoBasica(10, 5)
fmt.Println("Essa é sua soma:", soma)
fmt.Println("Essa é  sua multiplicacao:", multiplicacao)
fmt.Println("Essa é sua subtracao", subtracao)
}
	