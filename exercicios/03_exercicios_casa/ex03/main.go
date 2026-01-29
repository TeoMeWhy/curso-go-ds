// Faça um programa que receba o raio de uma circunferência em centímetros. Retorne para o usuário qual é a área e perímetro desta circunferência no seguinte formato.

// Área:  x.xx
// Perímetro:  y.yy

package main

import (
	"fmt"
	"math"
)

func main() {

	raio := 0.0
	fmt.Print("Entre com o raio da circunferência: ")
	fmt.Scanf("%f", &raio)

	area := math.Pi * raio * raio
	perimetro := 2 * math.Pi * raio

	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)

}
