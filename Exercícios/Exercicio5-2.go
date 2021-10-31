package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	mapa := make(map[string]pessoa)

	mapa["Santos"] = pessoa{
		nome:      "Gabriel",
		sobrenome: "Santos",
		sabores:   []string{"flocos", "chocolate", "baunilha"},
	}

	mapa["Barbosa"] = pessoa{
		nome:      "Adriano",
		sobrenome: "Barbosa",
		sabores:   []string{"uva", "creme", "leite ninho"},
	}

	for _, valor := range mapa {
		fmt.Printf("Meu nome é %v e meu sobrenome é %v, meus sabores favoritos são: \n", valor.nome, valor.sobrenome)
		for _, valor2 := range valor.sabores {
			fmt.Print(valor2, "\n")
		}
	}
}
