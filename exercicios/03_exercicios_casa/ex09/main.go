// Solicite ao usuário o nome de uma fruta e exiba o preço correspondente
// e some ao total de uma compra.
// Quando o usuário inserir uma string vazia,
// encerre a compra e informe o valor total da compra.

package main

import "fmt"

func main() {

	frutas := map[string]float64{
		"Maçã":    1.50,
		"Banana":  2.75,
		"Goiaba":  2.15,
		"Abacaxi": 3.20,
		"Jaca":    5.80,
		"Uva":     1.90,
		"Pera":    1.25,
		"Laranja": 0.65,
		"Limão":   1.25,
	}

	total := 0.
	for {

		fruta := ""
		fmt.Print("Entre com uma fruta: ")
		fmt.Scanf("%s", &fruta)

		if fruta == "" {
			break
		}

		if valor, ok := frutas[fruta]; !ok {
			fmt.Print("Entre com um valor válido!\n")
			continue
		} else {
			total += valor
		}

	}

	fmt.Printf("O valor total da sua compra é: R$%.2f", total)
}
