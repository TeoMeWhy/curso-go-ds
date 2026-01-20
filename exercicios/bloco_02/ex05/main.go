// Faça um programa que exiba
//  o dobro de um número inserido pelo usuário.

package main

import "fmt"

func main() {

	var x float64

	fmt.Println("Entre com um número:")
	fmt.Scanf("%f", &x)

	res := 2 * x
	fmt.Printf("O dobro de %.2f é %.2f\n", x, res)

}
