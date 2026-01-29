// Faça um programa que receba um número em segundos,
//  converta esse número para horas, minuto e segundos.
// Exemplos:

// Entrada: 556
// Saída: 0:9:16

// Entrada: 140153
// Saída: 38:55:53

package main

import "fmt"

func main() {

	var segundos, minutos, horas int

	fmt.Print("Entre com a quantidade de segundos: ")
	fmt.Scanf("%d", &segundos)

	horas = segundos / 3600
	segundos = segundos % 3600

	minutos = segundos / 60
	segundos = segundos % 60

	fmt.Printf("%02d:%02d:%02d\n", horas, minutos, segundos)

}
