# Função como expressão

- Assim como slices, ints, floats, strings, structs, métodos, as funções também podem ser usadas como valores assim como qualquer coisa como valor.

```GO
package main

import (
	"fmt"
)

func main() {
	x := 365
	y := func(x int) {
		fmt.Println(x, "vezes 45356 é:")
		fmt.Println(x * 45356)
	}
	y(x)
    // Exibe o tipo da variável y
	fmt.Printf("%T", y)
}
OUTPUT:
365 vezes 45356 é:
16554940
func(int)
```

