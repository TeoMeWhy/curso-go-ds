package main

import "fmt"

func main() {

	var x [100]int // isso e um array

	fatia := x[:10]
	fmt.Printf("x = %T | fatia = %T\n", x, fatia)
	fmt.Println("Tamanho fatia:", len(fatia))

	y := []int{} // isso é uma fatia
	fmt.Printf("y = %T\n", y)
	fmt.Println("Tamanho y:", len(y))

	for i := 1; i <= 200; i++ {
		y = append(y, i)
		fmt.Println(len(y))
	}

	fmt.Println(y)
}
