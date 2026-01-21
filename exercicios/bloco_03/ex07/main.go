// Faça um programa que a pessoa entra com 3 notas
//  e verificamos se a média delas é suficiente para passar na prova.
//  Nota de corte 6.

package main

import "fmt"

func main() {

	var nota, soma float64

	fmt.Println("Entre com a nota 1:")
	fmt.Scanf("%f", &nota)
	soma += nota

	fmt.Println("Entre com a nota 2:")
	fmt.Scanf("%f", &nota)
	soma = soma + nota

	fmt.Println("Entre com a nota 3:")
	fmt.Scanf("%f", &nota)
	soma += nota

	media := soma / 3.0

	if media >= 6 {
		fmt.Printf("Você passou! Média = %.2f\n", media)
	} else {
		fmt.Printf("Vai ter que estudar mais! Média = %.2f\n", media)
	}

}
