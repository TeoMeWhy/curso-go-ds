// Faça um programa que vende uma garrafa de água:
// Se o cliente escolher água mineral natural, será cobrado R$1,50
// Se o cliente escolher água mineral com gás, será cobrado R$2,50
// Altere o programa anterior para considerar a quantidade de garrafas de água

package main

import "fmt"

func main() {

	var opcao, qtde int
	fmt.Println("Você quer Água Mineral Natural (1) ou Água Mineral com Gás (2): ")
	fmt.Scanf("%d", &opcao)

	fmt.Println("Quantas garrafas você quer? ")
	fmt.Scanf("%d", &qtde)

	txtTemplate := "O valor ficou: R$%.2f x %d = R$%.2f\n"

	switch opcao {
	case 1:
		fmt.Printf(txtTemplate, 1.5, qtde, 1.5*float64(qtde))

	case 2:
		fmt.Printf(txtTemplate, 2.5, qtde, 2.5*float64(qtde))

	default:
		fmt.Println("Entre com um valor válido (1,2)")
	}

}
