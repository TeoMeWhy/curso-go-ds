// Faça um programa que receba dois valores A e B.
// Faça a potência desses dois valores e retorne o resultado:

// a ^ b = z

package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64

	fmt.Print("Entre com o valor de A: ")
	fmt.Scanf("%f", &a)

	fmt.Print("Entre com o valor de B: ")
	fmt.Scanf("%f", &b)

	res := math.Pow(a, b)
	fmt.Printf("A ^ B = %.2f ^ %.2f = %.2f\n", a, b, res)

}
