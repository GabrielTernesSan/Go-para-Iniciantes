# Funções anônimas

- Diferente das funções normais ela não possui nome e fica dentro do codeblock que será usada.
- Normalmente são funções de propósito rápido e descartável
- Podem ter elementos variádicos, vários retornos, etc. é uma função normal

````Go
package main

import (
	"fmt"
)

func main() {
	x := 365
	func(x int) {
		fmt.Println(x, "vezes 45356 é:")
		fmt.Println(x * 45356)
	}(x) // Equivalente a qualquercoisa(x)
}

/*
func qualquercoisa(x int){
    fmt.Println(x * 45356)
}
*/

````

