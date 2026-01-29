// Escreva um programa que solicite ao usuário seu nome e depois seu sobrenome.

// Exiba ambos depois utilizando apenas uma variável.

package main

import "fmt"

func main() {

	var nome, sobrenome string

	fmt.Print("Entre com o seu nome: ")
	fmt.Scanf("%s", &nome)

	fmt.Print("Entre com o seu sobrenome: ")
	fmt.Scanf("%s", &sobrenome)

	nomeCompleto := nome + " " + sobrenome
	fmt.Println("Seu nome completo é:", nomeCompleto)

}
