package main

import "fmt"

func soma(a int, b int) int {
	res := a + b
	return res
}

func media(a, b int) float64 {
	total := soma(a, b)
	res := float64(total) / float64(2)
	return res
}

func main() {

	n1, n2 := 10, 20
	fmt.Println(n1, n2)

	total := soma(n1, n2)
	fmt.Println("Total:", total)

	resMedia := media(n1, n2)
	fmt.Println("Média:", resMedia)

}
