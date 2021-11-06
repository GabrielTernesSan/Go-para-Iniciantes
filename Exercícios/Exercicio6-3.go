package main

import "fmt"

func main() {
	defer fmt.Println("Eu sou o primeiro a ser exebido")
	fmt.Println("Eu sou o segundo a ser exebido")
}
