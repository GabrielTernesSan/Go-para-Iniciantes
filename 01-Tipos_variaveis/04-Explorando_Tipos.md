# Explorando Tipos

- Tipos em GO são estáticos, ou seja, se você declarar uma variável com tipo Int ela será Int até o final.

- Ao declarar uma variável para conter valores de um tipo essa variável só poderá conter valores desse tipo.

- O tipo pode ser deduzido pelo compilador:

  ```GO
  x := 10 // Int
  var y := "a tia do Batima" // String
  ```

- Ou pode ser declarado especificamente

  ```GO
  var w string = "isso é uma string"
  var z int = 15
  ```

  

- **Tipos de dados primitivos**: Disponíveis na linguagem nativamente como blocos blocos básicos.

  ```GO
  int
  string
  bool
  ```

- **Tipos de dados compostos**: São tipos compostos de tipos primitivos e criados pelo usuário.

  - O ato de definir, criar, estruturar tipos compostos chama-se composição.

  ```Go
  slice
  array
  struct 
  map
  ```

- Package Reflect

  - Você pode usar a função do typeOf do package reflect para descobrir o tipo da variável

    ``````go
    package main
    
    import (
    	"fmt"
    	"reflect"
    )
    
    func main() {
    	nome := "Gabriel"
    	fmt.Println("O tipo de nome é", reflect.TypeOf(nome))
    }
    ``````

    