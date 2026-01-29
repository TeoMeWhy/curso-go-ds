// Faça um programa que receba dois valores A e B.
// Faça a soma desses dois valores e retorne o resultado:

// Soma:  x.xx

package main

import "fmt"

func soma(a, b float64) float64 {
	return a + b
}

func main() {

	var a, b float64
	fmt.Print("Entre com o valor de A: ")
	fmt.Scanf("%f", &a)

	fmt.Print("Entre com o valor de B: ")
	fmt.Scanf("%f", &b)

	res := soma(a, b)
	fmt.Printf("A + B = %.2f\n", res)

}
