// Faça um programa que verifique se a pessoa pertence à família “calvo” ou “silva”.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var nome string
	fmt.Println("Entre com o seu nome completo:")
	fmt.Scanf("%s", &nome)

	if strings.Contains(nome, "Calvo") || strings.Contains(nome, "Silva") {
		fmt.Println("Você é Silva ou Calvo!")
	} else {
		fmt.Println("Não conheço sua família!")
	}

}
