package main

import "fmt"

func main() {
	amigos := map[string]string{
		"Dias_Mauricio":    "Jogar video game",
		"Machado_Dienefer": "Ver Netflix",
		"Barbosa_Adriano":  "Estudar",
	}
	fmt.Println(amigos)
	for chave, valor := range amigos {
		fmt.Println(chave, valor)
	}
}
