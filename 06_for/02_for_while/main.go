package main

import "fmt"

func main() {

	for i := 26; i <= 1000; i++ {
		if i%25 == 0 {
			fmt.Println("Fodase", i)
			fmt.Println("Encontrei um múltiplo de 25")
			break
		}
	}
}
