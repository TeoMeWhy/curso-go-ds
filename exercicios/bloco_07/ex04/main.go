// Escreva um programa com a função “quadrado”,
//  que recebe um número inteiro e eleva ele ao quadrado,
//  retornando o resultado.

package main

import "fmt"

func quadrado(x int) int {
	return x * x
}

func main() {

	var x int
	fmt.Print("Entre com um número inteiro: ")
	fmt.Scanf("%d", &x)

	q := quadrado(x)
	fmt.Printf("%d x %d = %d\n", x, x, q)
}
