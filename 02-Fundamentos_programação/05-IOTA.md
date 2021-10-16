# IOTA

- Em uma declaração de constantes, o identificador IOTA representa números sequenciais não tipados.

  ```GO
  package main
  
  import "fmt"
  
  const (
  	x = iota
  	y = iota
  	z = iota
  )
  
  func main() {
  
  	fmt.Println(x, y, z)
  }
  
  OUTPUT: 0, 1, 2
  ```

  

- Podemos descartar valores

  ````GO
  package main
  
  import "fmt"
  
  const (
  	a = iota
  	_ = iota
  	c = iota
  	x = iota
  	_ = iota
  	z = iota
  )
  
  func main() {
  
  	fmt.Println(a, c, x, z)
  }
  
  OUTPUT: 0, 2, 3, 5
  
  
  package main
  
  import "fmt"
  
  const (
  	a = iota * 10
  	_ 
  	c 
  	x 
  	_ 
  	z 
  )
  
  func main() {
  
  	fmt.Println(a, c, x, z)
  }
  
  OUTPUT: 0, 20, 30, 50
  ````

  