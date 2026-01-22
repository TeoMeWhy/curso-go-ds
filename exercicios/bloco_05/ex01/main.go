// Faça um programa que receba 4 notas,
// calcule a média, mínimo e máximo dessas notas.

package main

import (
	"fmt"
	"slices"
)

func main() {

	var notas [4]float64 // isso aqui é um array

	for i := 1; i <= 4; i++ {
		var nota float64
		fmt.Printf("Entre com a %da nota: ", i)
		fmt.Scanf("%f", &nota)
		notas[i-1] = nota
	}

	min := slices.Min(notas[:])
	max := slices.Max(notas[:])

	var total float64
	for _, v := range notas {
		total += v
	}

	fmt.Println("Média:", total/float64(len(notas)))
	fmt.Println("Mínimo:", min)
	fmt.Println("Máximo:", max)
}
