package main

import "fmt"

func main() {

	nome := "Teodoro Calvo"

	for _, v := range nome {

		letra := string(v)
		if letra == "l" {
			break
		}

		if letra == " " {
			continue
		}

		fmt.Println(letra)
	}

}
