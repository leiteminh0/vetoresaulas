package main

import (
	"fmt"
)

var saldo int

func getsaldo() {
	fmt.Println("digite o quanto quer depositar")
	var deposito int
	fmt.Scan(&deposito)
	saldo = saldo + deposito
	fmt.Println("seu saldo atual é", saldo)
}

func getsacar() {
	var sacar int
	fmt.Println("digite oque deseja sacar")
	fmt.Scan(&sacar)
	saldo = saldo - sacar

}

func main() {
	fmt.Println("digite 1 para depositar, e 2 para sacar")
	var opcao int
	fmt.Scan(&opcao)
	if opcao == 1 {
		getsaldo()
	} else {
		getsacar()
	}
}
