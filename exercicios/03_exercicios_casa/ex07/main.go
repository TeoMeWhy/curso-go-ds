// Faça um programa que receba um número.
// Verifique se o número informado é par ou ímpar. Exiba o resultado da seguinte maneira:

// 	O número x é impar
// ou
// 	O número x é par

package main

import "fmt"

func ehPar(a int) bool {
	return a%2 == 0
}

func main() {

	var x int
	fmt.Printf("Entre com um número: ")
	fmt.Scanf("%d", &x)

	if ehPar(x) {
		fmt.Printf("O número %d é par!\n", x)
	} else {
		fmt.Printf("O número %d não é par!\n", x)
	}

}
