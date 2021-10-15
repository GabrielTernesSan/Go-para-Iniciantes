package main

import (
	"fmt"
)

type sal int

var x sal

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
