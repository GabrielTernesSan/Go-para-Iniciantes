# Tipo Booleano

- O tipo Bool é um tipo binário, que só pode conter um dos dois valores: true e false.

- Sempre que você ver operadores relacionais (==,<=,>=,!=,<,>), o resultado da expressão será um valor booleano.

- Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações IF, fluxo de controle, etc.

  ```GO
  package main
  
  import "fmt"
  
  func main() {
  	var b bool // Valor zero do tipo Bool é "false"
  	fmt.Printf("%v\n%T", b, b)
  }
  ```

  