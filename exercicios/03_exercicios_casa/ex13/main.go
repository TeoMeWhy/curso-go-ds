// Faça um programa que receba um número e exiba seu fatorial.
// Use função recursiva para isso.

package main

import "fmt"

func fatorial(x int) int {
	if x < 2 {
		return 1
	}

	return x * fatorial(x-1)
}

func main() {

	x := 0
	fmt.Print("Entre com um valor: ")
	fmt.Scanf("%d", &x)

	res := fatorial(x)
	fmt.Printf("%d! = %d\n", x, res)

}
