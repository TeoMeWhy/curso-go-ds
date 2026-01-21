// Faça um programa que receba uma quantidade indefinida
// de valores correspondentes a “saldo em conta”,
// mas quando o usuário apertar “enter” sem digitar valor algum,
// o programa para de receber valores,
// e exibe a soma de todos os valores digitados anteriormente.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var soma float64

	for {

		var entrada string // padrão é ""
		fmt.Print("Entre com o valor: ")
		fmt.Scanf("%s", &entrada)

		if entrada == "" {
			break
		}

		valor, erro := strconv.ParseFloat(entrada, 64)
		if erro != nil {
			fmt.Println("Entre com um valor válido!")
			continue
		}

		soma += valor
	}

	fmt.Printf("Saldo total em conta: R$%.2f\n", soma)
}
