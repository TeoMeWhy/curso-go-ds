package main

import "fmt"

func fib(n uint) int {

	if n <= 1 {
		return int(n)
	} else {
		return fib(n-1) + fib(n-2)
	}

}

func main() {

	fmt.Println("fib(0) =", fib(0))
	fmt.Println("fib(1) =", fib(1))
	fmt.Println("fib(2) =", fib(2))
	fmt.Println("fib(3) =", fib(3))
	fmt.Println("fib(4) =", fib(4))
	fmt.Println("fib(5) =", fib(5))

}
