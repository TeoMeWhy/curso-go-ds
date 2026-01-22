// Faça um programa que receba 6 temperaturas.
// Remova a 1a e a última para calcular a média.
// Se a média for acima de 30 graus,
// exiba que houve um aumento da temperatura.

package main

import "fmt"

func main() {

	var temps [6]float64 // isso é uma array

	for i := 1; i <= 6; i++ {
		var temp float64
		fmt.Printf("Entre com a %d temperatura: ", i)
		fmt.Scanf("%f", &temp)
		temps[i-1] = temp
	}

	newTemps := temps[1 : len(temps)-1]

	var total float64
	for _, v := range newTemps {
		total += v
	}

	media := total / float64(len(newTemps))

	if media > 30 {
		fmt.Println("Houve um aumento de temperatura")
	} else {
		fmt.Println("Tudo dentro da normalidade")
	}

}
