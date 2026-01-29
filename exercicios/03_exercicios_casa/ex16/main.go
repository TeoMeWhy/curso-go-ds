// Faça um programa que receba um número.
// Este número corresponde a uma posição na sequência de Fibonacci: 1, 1, 2, 3, 5,...

// Exiba o número da sequência cuja posição foi informada:
// 	A posição x corresponde ao número y

package main

import "fmt"

func fib(pos int) []int {
	values := []int{1, 1}
	for i := 2; i <= pos; i++ {
		values = append(values, values[i-1]+values[i-2])
	}
	return values[:pos]
}

func main() {

	x := 0
	fmt.Print("Entre com a posição da sequencia fibonacci: ")
	fmt.Scanf("%d", &x)

	seq := fib(x)
	fmt.Println(seq)

}
