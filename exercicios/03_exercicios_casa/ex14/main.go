// Escreva um programa que solicite ao usuário um número
// e exiba a tabuada desse número de 1 a 10.

package main

import "fmt"

func PrintTabuada(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %02d = %03d\n", x, i, x*i)
	}
}

func main() {
	x := 0
	fmt.Print("Entre com um valor para calcular a tabuada: ")
	fmt.Scanf("%d", &x)
	PrintTabuada(x)
}
