package main

import "fmt"

func main() {
	misterio := struct {
		piloto    map[string]string
		colocacao int
		equipe    []string
	}{
		piloto: map[string]string{
			"nome": "Fernando",
		},
		colocacao: 6,
		equipe:    []string{"RedBull, Ferrari, Mercedes"},
	}
	fmt.Println(misterio)
}
