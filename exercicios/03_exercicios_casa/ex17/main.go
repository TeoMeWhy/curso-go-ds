// Considere a seguinte lista:
// [123, 435, 987, 1984, 2, 19, 423, -178, 320]

// Faça um programa que retorne a posição do menor e do maior valor encontrado:

// O maior valor está na posição x
// O menor valor está na posição y

package main

import (
	"fmt"
	"slices"
)

func main() {

	values := []int{123, 435, 987, 1984, 2, 19, 423, -178, 320}

	maximo := slices.Max(values)
	minimo := slices.Min(values)
	posMaximo := 0
	posMinimo := 0

	for i, v := range values {
		if maximo == v {
			posMaximo = i
		}
		if minimo == v {
			posMinimo = i
		}
	}

	fmt.Println("O maior valor está na posição ", posMaximo)
	fmt.Println("O menor valor está na posição ", posMinimo)

}
