# Callback

- Callback é passar uma função como argumento.

  ````GO
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	t := somenteImpares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
  	fmt.Println(t)
  }
  
  //Faz a soma do "slice"
  func soma(x ...int) int {
  	n := 0
  	for _, v := range x {
  		n += v
  	}
  	return n
  }
  
  // Coloca todas os números impares dentro da "slice"
  //Passa como parâmetro uma função (soma()) de uma slice de números impares
  func somenteImpares(f func(x ...int) int, y ...int) int {
  	var slice []int
  	for _, v := range y {
  		if v%2 != 0 {
              // Atribui à slice os números impares 
  			slice = append(slice, v)
  		}
  	}
      // total recebe o resultado da função soma() que tem como argumento a slice de impares
  	total := f(slice...)
  	return total
  }
  
  ````

  

