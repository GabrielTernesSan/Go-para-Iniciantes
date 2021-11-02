# Métodos

- Um método é uma função anexada a um tipo.

  ```GO
  package main
  
  import (
  	"fmt"
  )
  
  type pessoa struct {
  	nome string
  	idade int
  }
  
  func (p pessoa) oibomdia() {
  	fmt.Println(p.nome, "diz bom dia!")
  }
  
  func main() {
  
  	mauricio := pessoa{"Maurício", 30}
  	mauricio.oibomdia()
  	
  }
  
  // func (receiver) identifier(parameters) (returns) { code }
  /* 
  O método anexa uma função a um determinado tipo, neste caso ao tipo
  pessoa, a função oibomdia() não funcionará sozinha ou com algum outro
  parâmetro
  */
  ```

  

- Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
- Pode-se anexar uma função a um tipo utilizando seu receiver.
- Utilização: valor.método()
- Exemplo: o tipo "pessoa" pode ter um método oibomdia()