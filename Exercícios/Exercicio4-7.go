package main

import "fmt"

func main() {
	slice := [][]string{
		{
			"Gabriel",
			"Ternes",
			"Programar",
		},
		{
			"Julia",
			"Santos",
			"Jogar volei",
		},
		{
			"Jonas",
			"Fraga",
			"Taekowndo",
		},
	}
	for _, valor := range slice {
		fmt.Println(valor)
	}
}
