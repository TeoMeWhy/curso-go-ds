package main

import "fmt"

func main() {

	notas := map[string][]float64{}
	notas["teo"] = []float64{7.89, 4.75, 9.89}
	notas["nah"] = []float64{10, 9.76, 9.89, 10, 4.89}
	fmt.Println(notas)
	fmt.Println(notas["teo"])

	cursos := map[string][]string{
		"teo": {"ds", "estatistica", "python"},
		"nah": {"arquitetura"},
	}

	cursos["lara"] = []string{"arquitetura", "engenharia"}
	fmt.Println(cursos)
}
