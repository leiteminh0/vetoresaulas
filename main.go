package main

import (
	"fmt"
)

func main() {
	alunoIdade := make(map[string]int)
	alunoIdade["Bangchan"] = 15
	alunoIdade["Felix"] = 16
	alunoIdade["Han"] = 15

	fmt.Println("Idade do Bangchan:", alunoIdade["Bangchan"])

	notasVocais := map[string]float64{
		"Bangchan": 100,
		"Felix":    95,
		"Han":      95,
	}

	for k, v := range notasVocais {
		fmt.Printf("%s tirou a nota %.1f \n", k, v)
	}
}
