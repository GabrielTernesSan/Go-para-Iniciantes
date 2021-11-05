# Recursividade

- Uma função recursiva é aquela capaz de chamar ela mesma. Um exemplo, uma função capaz de calcular o fatorial de um número:

  ```Go
  package main
  
  import "fmt"
  
  func main() {
      fmt.Println(fatorial(6))
  }
  
  func fatorial(x int) int{
      if x == 1{
          return x
      }
      return x * fatorial(x -1)
  } 
  ```

  - Podemos fazer utilizando loops de repetição

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	fmt.Println(fatorial(6))
    	fmt.Println(loops(6))
    
    }
    
    func fatorial(x int) int {
    	if x == 1 {
    		return x
    	}
    	return x * fatorial(x-1)
    }
    
    func loops(x int) int {
    	total := 1
    	for x > 1 {
    		total *= x // total = total * x
    		x--
    	}
    	return total
    }
    
    ```

    