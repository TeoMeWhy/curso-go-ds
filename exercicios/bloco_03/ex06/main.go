// Faça um programa que verifique se o item que a pessoa
//  escolheu para comprar na loja
//  está na lista: laranja, cerveja, miojo, carvão, picanha.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var opcao string
	fmt.Println("Entre com o item:")
	fmt.Scanf("%s", &opcao)

	itens := "laranja, cerveja, miojo, carvão, picanha"
	if strings.Contains(itens, opcao) {

		fmt.Println("Ok! Conheço esse item!")

	} else {

		fmt.Println("Acho que esse item não está no meu mercado.")

	}

}
