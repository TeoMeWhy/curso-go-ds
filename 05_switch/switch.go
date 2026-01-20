package main

import "fmt"

func main() {

	// "1- banana"
	// "2 - maca"
	// "3 - pera"
	// "4 - morando"

	var opcao int
	fmt.Println("Entre com uma opção: (1,2,3,4)")
	fmt.Scanf("%d", &opcao)

	switch opcao {

	case 1:
		fmt.Println("Você escolheu banana!")
		fmt.Println("Adoro bananas!")

	case 2:
		fmt.Println("hum... maça! Gosto muito! É a minha fruta predileta!")

	case 3:
		fmt.Println("Peras... já não gosto tanto.")

	case 4:
		fmt.Println("É, morangos não é comigo!")

	default:
		fmt.Println("Entre com uma opção válida!")

	}

}
