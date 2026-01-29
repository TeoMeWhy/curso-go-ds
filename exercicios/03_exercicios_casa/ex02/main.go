// Escreva um programa que receba o nome e a idade de uma pessoa. Depois exiba a mensagem:

// “Olá fulano, bom saber que você tem x anos. Boas vindas!”

package main

import "fmt"

func main() {

	var nome string
	var idade int

	fmt.Print("Qual o seu nome? ")
	fmt.Scanf("%s", &nome)

	fmt.Printf("Qual a sua idade? ")
	fmt.Scanf("%d", &idade)

	msg := "\nOlá %s, bom saber que você tem %d anos.\nBoas vindas!\n\n"
	fmt.Printf(msg, nome, idade)

}
