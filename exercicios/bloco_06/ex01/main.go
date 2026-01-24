// Faça o programa de uma sorveteria, onde o usuário pode escolher:
// Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
// Sabor do sorvete: morango, creme, chocolate
// Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
// Apresente o valor a ser pago.

package main

import (
	"fmt"
	"os"
)

func main() {

	tiposSorvete := map[string]float64{
		"casquinha": 1.00,
		"cascão":    2.50,
		"cestinha":  4.00,
	}

	saboresSorvete := map[string]float64{
		"morango":   0.0,
		"creme":     0.0,
		"chocolate": 0.0,
	}

	coberturasSorvete := map[string]float64{
		"caramelo":  1.5,
		"morango":   1.5,
		"chocolate": 1.5,
	}

	items := map[string]map[string]float64{
		"tipos":      tiposSorvete,
		"sabores":    saboresSorvete,
		"coberturas": coberturasSorvete,
	}

	var tipo, sabor, cobertura string
	var total float64

	fmt.Printf("Escolha um tipo [casquinha/cascão/cestinha]: ")
	fmt.Scanf("%s", &tipo)

	if valor, ok := items["tipos"][tipo]; !ok {
		fmt.Println("Tipo inválido!")
		os.Exit(1)
	} else {
		total += valor
	}

	fmt.Printf("Escolha o sabor [morango/creme/chocolate]: ")
	fmt.Scanf("%s", &sabor)

	if valor, ok := items["sabores"][sabor]; !ok {
		fmt.Println("Sabor inválido!")
		os.Exit(1)
	} else {
		total += valor
	}

	fmt.Printf("Escolha um cobertura [caramelo/morango/chocolate]: ")
	fmt.Scanf("%s", &cobertura)

	if valor, ok := items["coberturas"][cobertura]; !ok {
		fmt.Println("Cobertura inválida!")
		os.Exit(1)
	} else {
		total += valor
	}

	fmt.Printf("\nTotal: R$%.2f\n\n", total)
}
