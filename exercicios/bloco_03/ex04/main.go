// Faça um programa que verifique se a pessoa pertence à família “calvo”.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var nome string
	fmt.Println("Entre com o seu nome completo:")
	fmt.Scanf("%s", &nome)

	if strings.Contains(nome, "Calvo") {

		fmt.Println("Você deve ser parente do Téo...")

	} else {

		fmt.Println("Você não é careca!")

	}

}
