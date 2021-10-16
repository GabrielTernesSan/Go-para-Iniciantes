# Constantes

- São valores imutáveis

- Podem ser tipadas ou  não:

  ```Go
  package main 
  
  import "fmt"
  
  const (
  	oi = "Bom dia!"
      tchau string = "Boa noite!"
  )
  ```

- As não tipadas só terão um tipo atribuído a elas quando forem usadas.

  ```GO
  package main
  
  import "fmt"
  
  const x = 10 // Não sabemos se é uint, int ou float.
  var y = 10 // Já aqui, ao ler essa linha o compilador deduz que é um int e define assim.
  
  var c int
  var d float64
  
  func main() {
  	c = x // x passa a ser int
  	c = y // y é int
  	d = x // x passa a ser float64
  	d = y // Dará erro pois y é int e d é float64
  }
  ```

  