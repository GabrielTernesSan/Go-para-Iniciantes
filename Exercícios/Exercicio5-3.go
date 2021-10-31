package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	Toro := caminhonete{veiculo{2, "preta"}, true}
	Cronos := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "Branco",
		},
		modeloLuxo: false,
	}
	fmt.Println(Toro)
	fmt.Println(Cronos)
	fmt.Println(Cronos.cor)
	fmt.Println(Toro.portas)
}
