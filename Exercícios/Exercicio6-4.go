package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

var nome pessoa

func main() {
	Gabriel := pessoa{
		nome:      "Gabriel",
		sobrenome: "Santos",
		idade:     21,
	}
	Gabriel.ola()
}

func (p pessoa) ola() {
	fmt.Println("Olá", p.nome, p.sobrenome, "prazer em te conhecer, vi aqui no seu perfil que você tem", p.idade, "anos")
	fmt.Println(p.nome)
}
