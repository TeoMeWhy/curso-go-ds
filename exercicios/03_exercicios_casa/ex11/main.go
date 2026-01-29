// Escreva um programa que receba uma lista de números do
// usuário e conte quantas vezes um número específico aparece na lista.
// Solicite ao usuário um número e exiba a contagem.

package main

import "fmt"

func main() {

	values := map[int]int{}

	for {
		var v int
		fmt.Print("Entre com um número para ser registrado: ")
		fmt.Scanf("%d", &v)

		if v == 0 {
			break
		}
		values[v]++ // isso aqui é o charme!!
	}

	busca := 0
	fmt.Print("Entre com um número para descobrir quantas vezes foi digitado: ")
	fmt.Scanf("%d", &busca)

	fmt.Printf("o número %d apareceu %d vezes.", busca, values[busca])

}
