# Structs

- Struct é um tipo de dados compostos que nos permite armazenar valores de tipos _diferentes_.

- Seu nome vem de "structure", ou estrutura.

- Declaração

  ````GO
  type Pessoa struct {
      Idade int
      Nome string
      PCD bool    
  }
  ````

- Acesso

  ```GO
  fmt.Println(Pessoa.Nome)
  ```

  Exemplo: 

  ````GO
  package main
  
  import (
  	"fmt"
  )
  
  type cliente struct {
  	nome      string
  	sobrenome string
  	fumante   bool
  }
  
  func main() {
  	//Existem duas maneiras diferentes de declarar um struct
  	/*A primeira é esta, esta é a mais indicada pois fica
  	  mais legível para quem está acessando seu código*/
  	cliente1 := cliente{
  		nome:      "João",
  		sobrenome: "da Silva",
  		fumante:   false,
  	}
  
  	/*Esta é a segunda, desta maneira basta você declarar os dados na ordem
  	  que você definiu*/
  	cliente2 := cliente{"Joana", "Pereira", true}
  
  	fmt.Println(cliente1.nome)
  	fmt.Println(cliente2.fumante)
  	fmt.Println(cliente1)
  	fmt.Println(cliente2)
  }
  OUTPUT:
  João
  true
  {João da Silva false}
  {Joana Pereira true}
  ````

  