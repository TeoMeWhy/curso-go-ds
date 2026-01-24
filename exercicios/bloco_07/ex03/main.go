// Faça um programa que calcule a média de uma quantidade indefinida
// de valores, usando uma função para isso.

package main

import (
	"fmt"
	"strconv"
)

func soma(values ...float64) (res float64) {
	for _, v := range values {
		res += v
	}
	return
}

func media(values ...float64) float64 {

	if len(values) == 0 {
		return 0.0
	}

	return soma(values...) / float64(len(values))
}

func getValores() []float64 {

	values := []float64{}
	for {

		var valorStr string
		fmt.Printf("Entre com o valor: ")
		fmt.Scanf("%s", &valorStr)

		if valorStr == "" {
			return values
		}

		valor, err := strconv.ParseFloat(valorStr, 64)
		if err != nil {
			fmt.Println("Entre com um valor válido!")
			continue
		}

		values = append(values, valor)
	}
}

func main() {
	valores := getValores()
	mediaValor := media(valores...)
	fmt.Printf("Média: %.2f\n", mediaValor)
}
