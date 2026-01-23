// Faça um programa que informe se uma número é par ou ímpar
//  com uma função que se chama “EhPar”,
// retornando um booleano.

package main

import "fmt"

func EhPar(a int) bool {
	return a%2 == 0
}

func main() {

	var numero int
	fmt.Printf("Entre com um número: ")
	fmt.Scanf("%d", &numero)

	if EhPar(numero) {
		fmt.Println("É par!!")
	} else {
		fmt.Println("É ímpar!")
	}

}
