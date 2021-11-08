# Quando usar?

- Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
  - Não queremos passar grandes volumes de dados pra lá e pra cá
  - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)

````Go
package main

import (
	"fmt"
)

func main() {

	x := 11
	
	/*
	esta_recebe_o_valor(x) 
	Sendo assim, o valor é copiado e ao executar o x continua sendo 11 e a cópia
    dentro da função é incrementada.
	*/
    
    /*
    Já neste caso, o argumento passado é o endereço da variável x. A função recebe
    como parâmento um ponteiro de int, ou seja, o endereço passado, é feito o 
    incremento no valor do endereço recebido. 
    */
	esta_recebe_um_ponteiro(&x)
	
	fmt.Println(x)

}

func esta_recebe_o_valor(x int) {
	x++
	fmt.Println("Na função:", x)

}

func esta_recebe_um_ponteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
````

