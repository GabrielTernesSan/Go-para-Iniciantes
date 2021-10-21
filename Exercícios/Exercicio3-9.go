package main

import "fmt"

func main() {
	esporteFavorito := "Boxe"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Bora jogar uma bolinha?")
	case "Basquete":
		fmt.Println("Você deve ser alto...")
	case "Taekwondo":
		fmt.Println("Daora, gosta de kpop?:)")
	case "Boxe":
		fmt.Println("Esse ai é dos meus, Jab, Direto, upper!")
	}
}
