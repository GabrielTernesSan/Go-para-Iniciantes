# Condicionais

- A declaração If

  - If, else

  - If, else if, else

  - If, else if, else if, else

    ```GO
    package main
    
    import "fmt"
    
    func main(){
        x := 10
        if x < 10{
            fmt.Print("Olá mundo!")
        }
    }
    ```

- Switch

  ```GO
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	quemestanoescritoriohoje := "Mauricio"
  	
  	switch quemestanoescritoriohoje {
  		case "Mauricio":
  			fmt.Println("Boa Mauricio, pode colocar o café para passar ;)")
  		case "Gabriel":
  			fmt.Println("Boa Gabriel, pode colocar o café para passar ;)")
  		case "Adriano":
  			fmt.Println("Boa Adriano, pode colocar o café para passar ;)")
  	}
  	
  }
  OUTPUT:
  Boa Mauricio, pode colocar o café para passar ;)
  ```

  

  - Pode avaliar uma expressão

    - Switch statement == case (value)

    - Default switch statement == true (value)

    - ```GO
      func main() {
      	x := 4
      	
      	switch true{ //default
      		case (x == 4), (x == 8):
      			fmt.Println("É 4 ou 8")
      		case (x > 12):
      			fmt.Println("chis é maior que 10")
      		default:
      			// isso é o padrão se não houver outra opção
      	}
      	
      }
      OUTPUT:
      É 4 ou 8
      ```

  - Não há *Fallthrough*[^*] por padrão

  - Criando Fallthrough:

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	quemestanoescritoriohoje := "Mauricio"
    	
    	switch quemestanoescritoriohoje {
    		case "Mauricio":
    			fmt.Println("Boa Mauricio, Chegou cedo :)")
    			fallthrough
    		case "Gabriel":
    			fmt.Println("Mas como o Gabriel chegou antes, pode colocar o café para passar ;)")
    		case "Adriano":
    			fmt.Println("Boa Adriano, pode colocar o café para passar ;)")
    	}
    	
    }
    OUTPUT: 
    Boa Mauricio, Chegou cedo :)
    Mas como o Gabriel chegou antes, pode colocar o café para passar ;)
    ```

  - Cases compostos

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	quemestanoescritoriohoje := "Mauricio"
    	
    	switch quemestanoescritoriohoje {
    		case "Mauricio", "Brenda":
    			fmt.Println("Quem está trabalhando é a equipe 1")
    		case "Gabriel", "Bella":
    			fmt.Println("Quem está trabalhando é a equipe 2")
    		default:
    			fmt.Println("Cadê a equipe?! A balada estava boa ;-;")
    	}
    	
    }
    
    OUTPUT: 
    Quem está trabalhando é a equipe 1
    
    
    
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	x := 4
    	
    	switch {
    		case (x == 4), (x == 8):
    			fmt.Println("É 4 ou 8")
    		case (x > 12):
    			fmt.Println("chis é maior que 10")
    		default:
    			// isso é o padrão se não houver outra opção
    	}
    	
    }
    OUTPUT:
    chis é maior que 12
    ```

    

[^*]: "Se esse aqui está certo o próximo também", ou seja, se o case atual for true o próximo também será.

