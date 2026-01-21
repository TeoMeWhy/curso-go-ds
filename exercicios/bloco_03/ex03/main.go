// Faça o programa de uma sorveteria, onde o usuário pode escolher:
// Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
// Sabor do sorvete: morango, creme, chocolate
// Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
// Apresente o valor a ser pago

package main

import "fmt"

func main() {

	var tipoSorvete, saborSorvete, cobertura string

	fmt.Println("Qual o tipo de sorvete? [casquinha/cascão/cestinha]")
	fmt.Scanf("%s", &tipoSorvete)

	fmt.Println("Qual o sabor? [morango/creme/chocolate]")
	fmt.Scanf("%s", &saborSorvete)

	fmt.Println("Qual cobertura? [caramelo/morango/chocolate/sem cobertura]")
	fmt.Scanf("%s", &cobertura)

	var valorTipo, valorCobertura float64

	switch tipoSorvete {
	case "casquinha":
		valorTipo = 1.00
	case "cascão":
		valorTipo = 2.5
	case "cestinha":
		valorTipo = 4
	default:
		fmt.Println("Opção inválida para tipo")
	}

	if cobertura == "caramelo" || cobertura == "morango" || cobertura == "chocolate" {
		valorCobertura = 1.5
	} else if cobertura == "sem cobertura" {
		valorCobertura = 0
	} else {
		fmt.Println("Opção inválida para cobertura")
	}

	valorTotal := valorTipo + valorCobertura
	fmt.Printf("O total da conta foi: R$%.2f + R$%.2f = R$%.2f\n", valorTipo, valorCobertura, valorTotal)

}
