// Faça um programa que receba um número e exiba seu fatorial. Use função para isso.

package main

import "fmt"

func fatorial(x int) int {
	result := 1
	for i := x; i >= 2; i-- {
		result *= i
	}
	return result
}

func main() {

	var x int
	fmt.Print("Entre com um número: ")
	fmt.Scanf("%d", &x)

	res := fatorial(x)
	fmt.Printf("%d! = %d\n", x, res)

}
