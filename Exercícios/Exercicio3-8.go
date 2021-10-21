package main

import "fmt"

func main() {
	x := 10
	switch {
	case x == 10:
		fmt.Println("chis é igual a déix")
	case x == 12:
		fmt.Println("chis é igual a doze")
	case x == 15:
		fmt.Println("chis é igual a quinze")
	case x == 20:
		fmt.Println("chis é igual a vinte")
	}
}
