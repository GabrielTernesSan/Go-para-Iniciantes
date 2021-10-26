package main

import "fmt"

var x [5]int

func main() {
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	for Índice, valor := range x {
		fmt.Printf("Índice: %v, valor: %v \n", Índice, valor)
	}
	fmt.Printf("O tipo do array é %T", x)
}
