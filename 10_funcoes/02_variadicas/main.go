package main

import "fmt"

func soma(init int, values ...int) int {

	total := init
	for _, v := range values {
		total += v
	}

	return total
}

func main() {

	var a, b, c int

	a = 10
	b = 30
	c = 40

	total := soma(a, b, c, 100, 200, 300, 500, 1000)
	fmt.Println(total)

	valor := []int{1, 2, 3, 4, 5, 6, 10, 20, 34, 50, 321, 421, 54}
	total = soma(0, valor...)
	fmt.Println(total)

	fmt.Println("Olá,", "mundo", "!")
	fmt.Printf("Olá, %s\n", "Téo")

}
