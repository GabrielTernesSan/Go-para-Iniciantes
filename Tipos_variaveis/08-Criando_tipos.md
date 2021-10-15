# Criando seu próprio tipo

- Grande parte dos aspectos mais avançados de GO dependem quase que exclusivamente de tipos.

  **Revisando**: tipos são fixos. Uma vez declarada uma variável com de um certo tipo, isso é imutável.	

```GO
package main

import "fmt"

type gopher int

var g gopher

func main() {
	fmt.Printf("%T", g)//Exibe o tipo da variável
}

Output: 
main.gopher
```

- Uma variável do tipo "gopher" não pode ser atribuída com o valor de tipo int, mesmo que este seja o tipo subjacente de "gopher".