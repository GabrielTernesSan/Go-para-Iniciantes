package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := retorna()(x, y)
	fmt.Println(z)
}

func retorna() func(int, int) int {
	return func(x int, y int) int {
		soma := x + y
		return soma
	}
}
