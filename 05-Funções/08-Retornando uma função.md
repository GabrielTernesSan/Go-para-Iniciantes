# Retornando uma função

- Pode-se usar uma função como retorno de uma função

  ```GO
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  
  	x := retornaumafuncao()
  
  	y := x(3)
  
  	fmt.Println(y)
  	
      // Retorna a diretamente o "y"
  	fmt.Println(retornaumafuncao()(3))
      
      /* OU
      	x := retornaumafuncao()(3)
  		fmt.Println(x)
      */
  
  }
  
  //Essa função vai retornar uma função que irá receber um int e retornar um int
  func retornaumafuncao() func(int) int {
      // x = 
  	return func(i int) int {
  		return i * 10
  	}
  }
  
  ```

  

