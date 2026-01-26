package main

import "fmt"

type Altura float64 // em metros
type Peso float64   // em kilos

type Celsius float64
type Fahrenheit float64

func IMC(h Altura, p Peso) float64 {
	return float64(p) / float64(h*h)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(float64(c)*9.0/5.0 + 32.0)
}

func main() {

	alturaTeo := Altura(1.80)
	pesoTeo := Peso(70.34)

	res := float64(alturaTeo) + float64(pesoTeo)
	fmt.Println("altura + peso =", res)

	imc := IMC(alturaTeo, pesoTeo)
	fmt.Printf("IMC do Téo: %f\n", imc)

	c := Celsius(30)
	// f := Fahrenheit(15)

	fNew := CtoF(c)
	fmt.Printf("%f C = %f F", c, fNew)

}
