package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "Gabriel",
		sobrenome: "Santos",
		sabores:   []string{"flocos", "chocolate", "baunilha"},
	}

	pessoa2 := pessoa{
		nome:      "Adriano",
		sobrenome: "Barbosa",
		sabores:   []string{"uva", "creme", "leite ninho"},
	}

	fmt.Println(pessoa1.nome, pessoa2.sobrenome)
	for _, valor := range pessoa2.sabores {
		fmt.Println("\t", valor)
	}
}
