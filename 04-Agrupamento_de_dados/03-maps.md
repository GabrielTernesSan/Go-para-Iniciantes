# Maps

- Um mapa é uma coleção não ordenada de pares chave-valor (mapas ás vezes também são chamados de arrays associativos, tabelas hash ou dicionários). Os mapas são usados para  buscar um valor de acordo com sua chave associada.

- O tipo mapa é representado pela palavra reservada **_map_**, seguida do tipo da chave entre colchetes e, por fim, o tipo do valor.

- **Maps não tem ordem**

  ```GO
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  
  	amigos := map[string]int{
  		"alfredo": 5551234,
  		"joana":   9996674,
  	}
  	//Exibe o map inteiro
  	fmt.Println(amigos)
      //Exibe o valor associada a chave "joana"
  	fmt.Println(amigos["joana"])
  
      //Adiciona ao map o par chave-valor "gopher"
  	amigos["gopher"] = 444444
  
      //Exibe o map com o "gopher já incluido"
  	fmt.Println(amigos)
      
      //Exibe o valor associada a chave "gopher"
  	fmt.Println(amigos["gopher"], "\n\n")
  }
  OUTPUT:
  map[alfredo:5551234 joana:9996674]
  9996674
  map[alfredo:5551234 gopher:444444 joana:9996674]
  444444
  ```

- Se há um elemento que não exista dentro do Map, em GO, ele devolverá o valor zero do tipo do valor (que, para strings, é a string vazia). Embora pudéssemos verificar se o valor devolvido é o valor zero em uma condição, GO oferece uma opção melhor: o comma ok.

  - O comma ok basicamente diferencia o valor "0" em si do valor zero retornado pelo _map_ quando o elemento não existe.

    ```GO
    func main() {
    
    	amigos := map[string]int{
    		"alfredo": 5551234,
    		"joana":   9996674,
    	}
    
    	
    	fmt.Println(amigos)
    	fmt.Println(amigos["joana"])
    
    	amigos["gopher"] = 444444
    
    	fmt.Println(amigos)
    	fmt.Print(amigos["gopher"], "\n\n")
    
    
    	// será receberá o valor associado á chave
        // A variável ok é um boolean
        // ok = true; ele irá imprimir o valor associado a chave "fantasma"
        /* ok = false; ele irá exibir uma mensagem "não tem!" o que indica que 
        aquela chave não está dentro do map e seu valor é 0*/
    	if será, ok := amigos["fantasma"]; !ok {
    		fmt.Println("não tem!")
    	} else {
    		fmt.Println(será)
    	}
    
    }
    OUTPUT:
    map[alfredo:5551234 joana:9996674]
    9996674
    map[alfredo:5551234 gopher:444444 joana:9996674]
    444444
    
    não tem!
    
    /* caso a variável "ok" for igual a true ele irá retornar o valor
    correspondente a chave*/
    if será, ok := amigos["joana"]; !ok {
    		fmt.Println("não tem!")
    	} else {
    		fmt.Println(será)
    	}
    OUTPUT:
    map[alfredo:5551234 joana:9996674]
    9996674
    map[alfredo:5551234 gopher:444444 joana:9996674]
    444444
    
    9996674
    ```

  - Range e Delete em mapas

    - Range

      ```GO
      package main
      
      import (
      	"fmt"
      )
      
      func main() {
      
      	qualquercoisa := map[int]string{
      		123: "muito legal",
      		98:  "menos legal um pouquinho",
      		983: "esse é massa",
      		19:  "idade de ir pra festa",
      	}
      
      	//Exibe o mapa completo
      	fmt.Println(qualquercoisa)
      	
      	//Vai percorrer o mapa inteiro
      	for chave, valor := range qualquercoisa {
      		fmt.Println(chave, valor)
      	}
      
      }
      OUTPUT:
      //Perceba que não está na ordem
      map[19:idade de ir pra festa 98:menos legal um pouquinho 123:muito legal 983:esse é massa]
      
      123 muito legal
      98 menos legal um pouquinho
      983 esse é massa
      19 idade de ir pra festa
      ```

    - Delete

      ```GO
      package main
      
      import (
      	"fmt"
      )
      
      func main() {
      
      	qualquercoisa := map[int]string{
      		123: "muito legal",
      		98:  "menos legal um pouquinho",
      		983: "esse é massa",
      		19:  "idade de ir pra festa",
      	}
      
      	fmt.Println(qualquercoisa)
      
          //Basta usar a função delete(map, elemento)
      	delete(qualquercoisa, 123)
      
      	fmt.Println(qualquercoisa)
      
      }
      OUTPUT:
      map[19:idade de ir pra festa 98:menos legal um pouquinho 123:muito legal 983:esse é massa]
      map[19:idade de ir pra festa 98:menos legal um pouquinho 983:esse é massa]
      ```

      
