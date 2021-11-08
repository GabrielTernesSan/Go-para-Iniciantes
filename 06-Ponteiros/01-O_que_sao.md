# Ponteiros

- Todos os valores ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um pointeiro se refere a esse endereço.
- Notações:
  - &variável mostra o endereço de uma variável
    - %T: variável vs. &variável
  - *variável faz de-reference, mostra o valor que consta nesse endereço
  - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador

````go
package main

import (
	"fmt"
)

func main() {

	x := 0

    //Mostra o endereço na memória da variável x
	y := &x

    //Exibe as variáveis
	fmt.Print(x, y)

    /*
    Atribui 1 ao valor armazenado no endereço da variável x, já que a notação *y 
    mostra o valor que consta no endereço e não a variável diretamente
    */
	*y++

    //Exibe o valor que consta no endereço
	fmt.Println(*y)
    //Exibe o tipo das variáveis x(int) e y(*int ponteiro que aponta para um int)
	fmt.Printf("%T, %T\n", x, y)
     //Exibe as variáveis após o incremento
	fmt.Print(x, y)
}

````

