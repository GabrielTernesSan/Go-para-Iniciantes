# Array

- Um array é uma sequência numerada de elementos de um único tipo com um tamanho fixo. Em GO, eles tem o seguinte aspecto:

  ```go
  var x [5]int // array de 5 posições (0-4)
  ```

- Arrays em Go são uns dos fundamentos da linguagem, não são muito utilizados. Arrays são úteis para planejar o layout da memória e servem como fundação para slices. Recomenda-se que se utilize slices ao invés de arrays.

- Atribui-se valores a suas posições com x[i] = y.

- Para ver seu tamanho usa-se: len(x).

  ```GO
  package main
  
  import (
  	"fmt"
  )
  
  var x [5]int
  var y [6]int
  
  func main() {
  	x[0] = 1
  	x[1] = 10
  	fmt.Println(x[0], x[1])
  	fmt.Println(x)		// exibe todo o array
  	fmt.Printf("%T\n", x) // retorna o tipo do array ([5]int)
  	fmt.Printf("%T\n", y) // retorna o tipo do array ([6]int)
  	// tipos diferentes e incompatíveis
  	fmt.Println(len(x)) // retorna o tamanho do array
  
  }
  ```

  