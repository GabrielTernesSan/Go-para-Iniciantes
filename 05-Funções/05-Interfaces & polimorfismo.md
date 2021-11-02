# Interfaces & polimorfismo

- Em Go, valores podem ter mais que um tipo.

- Uma interface permite que um valor tenha mais que um tipo.

- Declaração: keyword identifier type → type x interface

- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.

- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.

- Esse tipo será o seu tipo *e também* o tipo da interface.

  

  Exemplo:

  - Os tipos *profissão1* e *profissão2* contem o tipo *pessoa*

  - Cada um tem seu método *oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()*

  - Implementam a interface *gente*

  - Ambos podem acessar o função *serhumano()* que chama o método *oibomdia()* de cada *gente*

  - Também podemos no método *serhumano()* tomar ações diferentes dependendo do tipo: switch pessoa.(type) { 

    case profissão1: 

    fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...]

     }*

  

  ````Go
  package main
  
  import (
  	"fmt"
  )
  
  // Tipo pessoa será comum aos dois tipos "profissão"
  type pessoa struct {
  	nome      string
  	sobrenome string
  	idade     int
  }
  
  // Profissão 1, contém "pessoa"
  type arquiteto struct {
  	pessoa
  	especialização  string
  	obrasconcluidas int
  }
  
  // Profissão 2, contém "pessoa"
  type dentista struct {
  	pessoa
  	especialização   string
  	dentesarrancados int
  }
  
  // A função "oibomdia()" é anexada ao tipo arquiteto
  func (a arquiteto) oibomdia() {
  	fmt.Println("Oi, bom dia! Meu nome é", a.nome, "e eu já fiz", a.obrasconcluidas, "prédios.")
  }
  
  // A função "oibomdia()" é anexada ao tipo dentista
  func (d dentista) oibomdia() {
  	fmt.Println("Oi, bom dia! Meu nome é", d.nome, "e eu já arranquei", d.dentesarrancados, "dentes.")
  }
  
  //Implementação da interface
  type gente interface {
      // Define o método para implementar a interface
  	oibomdia()
  }
  
  // A função "serhumano" chama o método "oibomdia()" de cada tipo (dentista e arquiteto)
  func serhumano(g gente) {
  
  	switch g.(type) {
        
      /* 
      Caso o tipo seja "arquiteto" será chamado o método "oibomdia()" com
      o argumento arquiteto. O que exibirá a mensagem no case + a mensagem da
      função "oibomdia":
      Este arquiteto é especializado em Galpão de fazenda e já fez 20 obras.
      Ele diz: Oi, bom dia! Meu nome é Paulo e eu já fiz 20 prédios.
      */
      case arquiteto:
  		fmt.Println("Este arquiteto é especializado em", g.(arquiteto).especialização, "e já fez", g.(arquiteto).obrasconcluidas, "obras. Ele diz:")
          
      /* 
      Caso o tipo seja "dentista" será chamado o método "oibomdia()" com
      o argumento dentista. O que exibirá a mensagem no case + a mensagem da
      função "oibomdia":
      Este dentista é especializado em Tortura e já arrancou 8748 dentes. 
      Ele diz: Oi, bom dia! Meu nome é Henrique e eu já arranquei 8748 dentes.
      */    
  	case dentista:
  		fmt.Println("Este dentista é especializado em", g.(dentista).especialização, "e já arrancou", g.(dentista).dentesarrancados, "dentes. Ele diz:")
  	}
      // Chamada da função "oibomdia" com o argumento (type)
  	g.oibomdia()
  }
  
  func main() {
  
      // Define valores das chaves
  	pessoa1 := arquiteto{
  		pessoa: pessoa{
  			nome:      "Paulo",
  			sobrenome: "Prédio",
  			idade:     40,
  		},
  		especialização:  "Galpão de fazenda",
  		obrasconcluidas: 20,
  	}
  
  	pessoa2 := dentista{
  		pessoa: pessoa{
  			nome:      "Henrique",
  			sobrenome: "Cido",
  			idade:     50,
  		},
  		especialização:   "Tortura",
  		dentesarrancados: 8748,
  	}
  
  	/*
          fmt.Println(pessoa1, pessoa2)
          {{Paulo Prédio 40} Galpão de fazenda 20} 
          {{Henrique Cido 50} Tortura 8748}
      */
  
  	/*
          pessoa1.oibomdia()
          Oi, bom dia! Meu nome é Paulo e eu já fiz 20 prédios.
      */
      
  	/*
          pessoa2.oibomdia()
          Oi, bom dia! Meu nome é Henrique e eu já arranquei 8748 dentes.
  	*/
      
  	/*
  		serhumano(pessoa1)
  		Este arquiteto é especializado em Galpão de fazenda e já fez 20
          obras. Ele diz: Oi, bom dia! Meu nome é Paulo e eu já fiz 20 prédios.
  	*/
      
  	/*
  		serhumano(pessoa2)
  		Este dentista é especializado em Tortura e já arrancou 8748 dentes.
           Ele diz: Oi, bom dia! Meu nome é Henrique e eu já arranquei 
           8748 dentes.
  	*/
  }
  ````

