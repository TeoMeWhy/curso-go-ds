// Faça um programa que de bom dia,
// pergunta o nome da pessoa e
// responde que é um prazer conhecer ela, citando o nome da pessoa.

package main

import "fmt"

func main() {

	fmt.Println("Bom dia! Qual é o seu nome?")

	var nome string
	fmt.Scanf("%s", &nome)
	fmt.Printf("É um prazer te conhecer, %s.\n", nome)

}
