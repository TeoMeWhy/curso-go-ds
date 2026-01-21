package main

import "fmt"

func main() {

	nome := "Teo Calvo"

	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
	}

}
