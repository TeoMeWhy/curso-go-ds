package main

import "fmt"

func main() {

	idades := make(map[string]int)
	idades["teo"] = 33
	idades["nah"] = 34
	fmt.Println("Idades:", idades)

	alturas := map[string]float64{}
	alturas["teo"] = 1.82
	fmt.Println("Alturas:", alturas)

	alturaTeo := alturas["teo"]
	fmt.Println("Altura Téo:", alturaTeo)

	alturas["nah"] = 1.70
	alturaNah, ok := alturas["nah"]
	if ok {
		fmt.Println("Altura Nah:", alturaNah)
	} else {
		fmt.Println("Não encontrei!")
	}

	if valor, ok := alturas["lara"]; ok {
		fmt.Println("A altura da lara é:", valor)
	} else {
		fmt.Println("Não encontrei a altura da Lara!")
	}

	fmt.Println(ok)

}
