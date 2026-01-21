package main

import "fmt"

func main() {

	nome := "Teo Calvo"

	for _, v := range nome {
		fmt.Println(string(v))
	}

}
