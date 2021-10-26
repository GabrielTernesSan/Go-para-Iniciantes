package main

import "fmt"

func main() {
	slice := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	for Índice, valor := range slice {
		fmt.Println("No Índice,", Índice, "temos o valor", valor)
	}
	fmt.Printf("A slice é o tipo %T", slice)
}
