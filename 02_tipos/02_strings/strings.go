package main

import "fmt"

func main() {

	fmt.Println("O que você quiser!!")

	fmt.Println(`Aqui a gente pode escrever
Múltiplas linhas
de
código!`)

	fmt.Println("Teodoro tem ", len("teodoro"), "letras")

	fmt.Println("Teodoro"[0]) // isso é um byte (int8)

	fmt.Println(string("Teodoro"[0]))
	fmt.Println(string("Teodoro"[1]))
	fmt.Println(string("Teodoro"[2]))

}
