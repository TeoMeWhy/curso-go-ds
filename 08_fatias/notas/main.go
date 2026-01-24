package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var notas []float64
	notas := []float64{}

	for {

		var inputTxt string
		fmt.Printf("Ente com a nota: ")
		fmt.Scanf("%s", &inputTxt)

		if inputTxt == "" {
			break
		}

		nota, err := strconv.ParseFloat(inputTxt, 64)
		if err != nil {
			fmt.Println("Entre com uma nota válida!")
			continue // para solicitar nova entrada!
		}

		notas = append(notas, nota)
	}

	fmt.Println(notas)
}
