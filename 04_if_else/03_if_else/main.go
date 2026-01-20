package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Entre com a sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 66 {

		fmt.Println("Olha só heim...")
		fmt.Println("Manere com a bebedeira!")

	} else if idade >= 18 {

		fmt.Println("Perfeito!")
		fmt.Println("Beba a vontade!")

	} else {

		fmt.Println("Vishhhh")
		fmt.Println("Melhor ir para casa tomar suco")

	}

}
