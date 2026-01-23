package main

import "fmt"

func primeira() {
	fmt.Println("Eu sou chamada primeira")
}

func segunda() {
	fmt.Println("Eu sou chamada segunda")
}

func terceira() {
	fmt.Println("Eu sou chamada terceira")
}

func ultima() {
	fmt.Println("Eu sou chamada ultima")
}

func main() {
	defer ultima()   // última
	defer terceira() // penútima

	primeira()
	segunda()

}
