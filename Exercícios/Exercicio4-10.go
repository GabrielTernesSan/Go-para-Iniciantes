package main

import "fmt"

func main() {
	amigos := map[string]string{
		"Dias_Mauricio":    "Jogar video game",
		"Machado_Dienefer": "Ver Netflix",
		"Barbosa_Adriano":  "Estudar",
	}
	amigos["Costa_Brenda"] = "namorar"

	delete(amigos, "Machado_Dienefer")

	for chave, valor := range amigos {
		fmt.Println(chave, valor)
	}
}
