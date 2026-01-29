// Escreva um programa que receba o nome de uma pessoa e faça uma saudação.

// “Olá fulano! Boas vindas!”

package main

import "fmt"

func main() {

	var nome string
	fmt.Print("Entre com o seu nome: ")
	fmt.Scanf("%s", &nome)

	fmt.Printf("Olá, %s! Boas vindas!\n", nome)

}
