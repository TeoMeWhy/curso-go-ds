// Escreva um programa que exiba os números de 1 a 100.
// Caso o número seja divisível por 3, exiba “Fizz” no seu lugar,
// e para múltiplos de 5 exiba “Buzz”.
// Caso seja divisível por ambos, exiba “FizzBuzz”.

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		var msg string
		if i%3 == 0 {
			msg += "Fizz"
		}

		if i%5 == 0 {
			msg += "Buzz"
		}

		if msg == "" {
			fmt.Println(i)
			continue
		}

		fmt.Println(msg)
	}
}
