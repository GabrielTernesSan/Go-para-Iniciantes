package main

import "fmt"

func main() {
	a := s()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func s() func() int {
	x := 0
	y := 1
	z := 2
	return func() int {
		x++
		y++
		z++
		t := x + y + z
		return t
	}
}
