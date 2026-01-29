// Faça um programa que receba um número.
// Verifique se este número é primo ou não, e retorne o resultado:

// 	O número x é primo
// ou
// 	O número x não é primo

package main

import "fmt"

func EhPrimo(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	x := 0
	fmt.Print("Entre com um número para verificar se é primo: ")
	fmt.Scanf("%d", &x)

	if EhPrimo(x) {
		fmt.Printf("%d é primo!\n", x)
	} else {
		fmt.Printf("%d não é primo!\n", x)
	}

}
