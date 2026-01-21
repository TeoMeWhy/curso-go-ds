// Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra

package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra string
	fmt.Println("Entre com uma palavra:")
	fmt.Scanf("%s", &palavra)

	// para garantir que vai estar tudo em caixa baixa!!
	palavraLower := strings.ToLower(palavra)

	contador := 0
	for _, letraByte := range palavraLower {

		if string(letraByte) == "a" {
			contador++
		}

	}

	fmt.Printf("A palavra '%s' tem %d As\n", palavra, contador)
}
