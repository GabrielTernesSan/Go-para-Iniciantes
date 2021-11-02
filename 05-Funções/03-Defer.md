# Defer

- Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!

- Uma declaração defer chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.

  ````Go
  package main
  
  import "fmt"
  
  func main() {
  	defer fmt.Println("Primeiro ;)")
  	fmt.Println("Segundo")
  
  /*
  OUTPUT: 
  Segundo
  Primeiro ;)
  */
  
      // Entrou primeiro no defer, então será o último a ser exibido
  	defer fmt.Println("1") 
  	defer fmt.Println("2")
  	fmt.Println("3")
  	fmt.Println("4")	
  /*
  OUTPUT: 
  3
  4
  2
  1
  */
  }
  
  ````

  

- Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.

- Sempre usamos para fechar um arquivo após abri-lo.

   ````GO
   f, _ := os.Open(filename)
   defer f.Close()
   ````

  

