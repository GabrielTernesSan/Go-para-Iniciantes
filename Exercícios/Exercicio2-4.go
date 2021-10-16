package main

import (
	"fmt"
)

func main() {
	x := 15
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\n%b\n%#x", y, y, y)
}
