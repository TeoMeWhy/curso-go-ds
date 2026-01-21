// Faça um programa que vende uma garrafa de água:
// Se o cliente escolher água mineral natural, será cobrado R$1,50
// Se o cliente escolher água mineral com gás, será cobrado R$2,50

package main

import "fmt"

func main() {

	var opcao int
	fmt.Println("Você quer Água Mineral Natural (1) ou Água Mineral com Gás (2): ")
	fmt.Scanf("%d", &opcao)

	txtTemplate := "O valor ficou: R$%.2f\n"

	switch opcao {
	case 1:
		fmt.Printf(txtTemplate, 1.5)

	case 2:
		fmt.Printf(txtTemplate, 2.5)

	default:
		fmt.Println("Entre com um valor válido (1,2)")
	}

}
