package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ { // Código decimal das letras maiúsculas 65 = A ... Z = 90
		fmt.Println(i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
