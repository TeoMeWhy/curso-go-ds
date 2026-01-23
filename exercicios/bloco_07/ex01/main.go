// Faça um programa que use uma função para responder
//  a saudação de um usuário: “Olá, fulano! Boas vindas!”.
// Use funções para isso.

package main

import "fmt"

func olaFulano(nome string) {
	fmt.Printf("Olá, %s! Boas vindas!\n", nome)
}

func main() {

	var nome string

	fmt.Printf("Entre com o seu nome: ")
	fmt.Scanf("%s", &nome)
	olaFulano(nome)

	fmt.Printf("Entre com o seu nome: ")
	fmt.Scanf("%s", &nome)
	olaFulano(nome)

}
