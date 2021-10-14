# A palavra-chave Var

- Variáveis declaradas em um codeblock é undefined em outro.

```GO
package main

import "fmt"

func main() {
	y := 10
	exibe(y)
}

func exibe(x int) {
	fmt.Println(y) // undefined: y, dentro da função "exibe" o y não existe. Não está dentro do escopo.
	fmt.Println(x)
}

```

- Para variáveis com uma abrangência maior, package level scope, utilizamos "Var".

```Go
package main

import "fmt"

// Neste caso funcionará pois a variável y está em um escopo global, todos podem vê-la.

var y int = 10 

func main() {
	exibe(y)
}

func exibe(x int) {
	fmt.Println(y) 
	fmt.Println(x)
}
```

- Funciona em qualquer lugar

  ```GO
  package main
  
  import "fmt"
  
  var y int = 10
  
  func main() {
  	var x int = 15
  	fmt.Println(x)
  	fmt.Println(y)
  }
  ```

  
