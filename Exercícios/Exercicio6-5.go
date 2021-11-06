package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	area := q.lado * q.lado
	fmt.Println("A área do quadrado é:", area)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	area := 2 * math.Pi * c.raio
	fmt.Println("A área do circulo é:", area)
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}

func main() {
	x := quadrado{
		lado: 15.0,
	}
	y := circulo{
		raio: 10.0,
	}
	info(x)
	info(y)
}
