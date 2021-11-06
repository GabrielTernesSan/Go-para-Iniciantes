package main

import (
	"fmt"
)

func main() {
	t := raizQuadrada(num, 15)
	fmt.Println(t)
}

func raizQuadrada(f func(int) int, z int) int {
	total := z * z
	return total
}

func num(x int) int {
	return x
}
