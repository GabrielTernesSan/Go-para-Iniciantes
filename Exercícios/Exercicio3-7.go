package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		if x%4 == 0 { // mostra os numeros divisiveis por 4
			fmt.Println(x)
		} else {
			fmt.Println("Não é divisivel")
		}
	}
}
