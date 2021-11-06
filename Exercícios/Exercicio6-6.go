package main

import "fmt"

func main() {
	x := "Lasanha"

	func(x string) {
		fmt.Print("Minha comida preferida é: ", x)
	}(x)

}
