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

- Slice embutidos

  - São structs dentro de strucs dentro de structs

  - Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) *e também* um competidor (nome, equipe, pontos).

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    type pessoa struct {
    	nome      string
    	sobrenome string
    	idade     int
    }
    
    type piloto struct {
    	pessoa
    	nome    string
    	equipe  string
    	pontos  int
    	tempo   float64
    	posicao int
    }
    
    func main() {
    
    	pessoa1 := pessoa{
    		nome:      "Gabriel",
    		sobrenome: "Santos",
    		idade:     21,
    	}
    	pessoa2 := piloto{
    		pessoa: pessoa{
    			nome:      "Max",
    			sobrenome: "Verstappen",
    			idade:     24,
    		},
    		pontos:  25,
    		tempo:   1.34,
    		posicao: 1,
    	}
    	pessoa3 := piloto{
    		pessoa: pessoa{
    			nome:      "Lewis",
    			sobrenome: "Hamilton",
    			idade:     36,
    		},
    		pontos:  19,
    		tempo:   1.35,
    		posicao: 2,
    	}
    
    	fmt.Println(pessoa1, pessoa2, pessoa3)
    }
    OUTPUT:
    {Gabriel Santos 21}{{Max Verstappen 24}25 1.34 1}{{Lewis Hamilton 36} 19 1.35 2}
    
    ```

    

