// Crie uma história simples.
// Adicione essa história em um programa.
// A cada parágrafo, a história deve aguardar o
//  usuário apertar “enter” para dar continuidade.

package main

import "fmt"

func main() {

	var ok string

	fmt.Println("Era uma vez, um mago muito poderoso")
	fmt.Scanf("%s", &ok)

	fmt.Println("Seu melhor amigo era uma marmota (castor).")
	fmt.Scanf("%s", &ok)

	fmt.Println("Juntos desbravaram o mundo da programação.")
	fmt.Scanf("%s", &ok)

	fmt.Println("Aventuras incríveis estavam lhes esperando.")
	fmt.Scanf("%s", &ok)

}
