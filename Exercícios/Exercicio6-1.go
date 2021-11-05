package main

import "fmt"

func main() {
	fmt.Println(Inteiro())
	fmt.Println(IntMaisString())
}

func Inteiro() int {
	x := 5
	return x
}

func IntMaisString() (string, int) {
	y := "Números de exercícios feitos até o momento:"
	z := 34
	return y, z
}
