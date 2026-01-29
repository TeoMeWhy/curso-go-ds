// Escreva um programa que solicite ao usuário uma palavra e
//  verifique se a palavra é um palíndromo
// (ou seja, é a mesma palavra quando lida de trás para frente).

package main

import (
	"fmt"
	"slices"
)

func EhPalindromo(palavra string) bool {

	palavraSlice := []string{}

	for _, letra := range palavra {
		palavraSlice = append(palavraSlice, string(letra))
	}

	palavraSliceInvert := palavraSlice[:]
	slices.Reverse(palavraSliceInvert)

	for i := range palavraSlice {
		if palavraSlice[i] != palavraSliceInvert[i] {
			return false
		}
	}
	return true
}

func main() {

	var palavra string
	fmt.Print("Entre com uma palavra: ")
	fmt.Scanf("%s", &palavra)
	if EhPalindromo(palavra) {
		fmt.Println("É palíndromo!")
	} else {
		fmt.Println("Não é palíndromo!")
	}

}
