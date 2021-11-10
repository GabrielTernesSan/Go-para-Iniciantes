package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func main() {
	Gabriel := pessoa{
		nome:      "João",
		sobrenome: "Guilherme",
	}
	fmt.Println(Gabriel)
	mudeMe(&Gabriel)
	fmt.Println(Gabriel)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Gabriel"
	p.sobrenome = "santos"
}
