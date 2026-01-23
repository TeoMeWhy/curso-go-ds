package main

import "fmt"

func addOne(a int) int {
	fmt.Println("Endereço de a:", &a)
	a++ // a = a + 1
	return a
}

func addOnePt(a *int) {
	fmt.Println("Endereço do ponteiro a:", a)
	fmt.Println("Valor do ponteiro de a:", *a)
	*a++
}

func main() {

	numero := 10
	fmt.Println("Endereço de 'numero'", &numero)

	valor := addOne(numero) // isso gera uma cópia,
	// não é o mesmo objeto que será usado no corpo da função

	addOnePt(&numero)
	fmt.Println("Número:", numero)
	fmt.Println("Valor:", valor)

}
