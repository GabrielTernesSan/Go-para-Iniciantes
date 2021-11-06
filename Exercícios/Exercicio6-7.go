package main

import "fmt"

func main() {
	x := 45
	y := func(x int) int {
		x = x * 2
		return x
	}
	fmt.Println(y(x))
}
