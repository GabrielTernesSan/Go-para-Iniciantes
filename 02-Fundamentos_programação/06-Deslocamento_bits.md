# Deslocamento de Bits

- Deslocamento de bits é quando deslocamos dígitos binários para a esquerda ou direita.

  ```GO
  package main
  
  import "fmt"
  
  func main() {
  
  	x := 1
  	y := x << 1 // deslocando para a esquerda
  	fmt.Printf("%b\n", x)
  	fmt.Printf("%b\n", y)
  }
  
  OUTPUT: 1 10
  
  ##
  
  package main
  
  import "fmt"
  
  func main() {
  
  	x := 2
  	y := x >> 1 // deslocando para a direita
  	fmt.Printf("%b\n", x)
  	fmt.Printf("%b\n", y)
  }
  
  OUTPUT: 10 1
  
  ##
  
  package main
  
  import (
  	"fmt"
  )
  
  const (
  	_  = iota             // 0
  	KB = 1 << (iota * 10) // 1 << (1 * 10)
  	MB = 1 << (iota * 10) // 1 << (2 * 10)
  	GB = 1 << (iota * 10) // 1 << (3 * 10)
  	TB = 1 << (iota * 10) // 1 << (4 * 10)
  )
  
  func main() {
  	fmt.Println("binary\t\t\t\tdecimal")
  	fmt.Printf("%b\t\t\t", KB)
  	fmt.Printf("%d\n", KB)
  	fmt.Printf("%b\t\t", MB)
  	fmt.Printf("%d\n", MB)
  	fmt.Printf("%b\t", GB)
  	fmt.Printf("%d\n", GB)
  }
  
  OUTPUT: 
  binary								decimal
  10000000000							1024
  100000000000000000000				1048576
  1000000000000000000000000000000		1073741824
  ```

  