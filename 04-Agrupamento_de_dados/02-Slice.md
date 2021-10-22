# Slice

- Uma slice (fatia) é um segmento de um array. Assim como so arrays, as fatias podem ser indexadas e têm tamanhos. De modo diferente dos arrays, o tamanho pode ser mudado.

- Uma slice agrupa valores de um único tipo.

  ```GO
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
      //Array de 5 posições fixo
  	array := [5]int{1, 2, 3, 4, 5}
  	fmt.Println(array)
      //Slice de 5 posições
  	slice := []int{1, 2, 3, 4, 5}
  	fmt.Println(slice)
  
  	slice2 := append(slice, 6) 
      // Ao contrário do Array o slice não possui um tamanho fixo, sendo assim, é possível adicionar valores.
  	fmt.Println(slice2)
  	
  	//Mudar o valor de um índice específico
  	fmt.Println(slice[3])
  	slice[3] = 348756
  	fmt.Println(slice[3])
  }
  OUTPUT: 
  [1 2 3 4 5]
  [1 2 3 4 5]
  [1 2 3 4 5 6]
  4
  348756
  ```

- Slice for range

  - Para ver todos os itens de uma slice utilizamos o loop for com range.

  - Range significa alcance, faixa, extensão.

  - For range:

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	slice := []string{"banana", "maçã", "jaca", "pêssego"}
    
    	//for índice, valor := range slice {fmt.Println("No índice", índice, "temos o valor:", valor)}
    
    	//slice[4] = "melancia"
    	slice = append(slice, "melancia")
    
        // A função range retorna o índice e o valor no índice referido
    	for índice, valor := range slice {
    		fmt.Printf("No índice", índice, "temos o valor", valor)
    	}
    }
    OUTPUT:
    No índice 0 temos o valor banana
    No índice 1 temos o valor maçã
    No índice 2 temos o valor jaca
    No índice 3 temos o valor pêssego
    No índice 4 temos o valor melancia
    
    // Podemos também retornar somente os valores
    for _, valor := range slice {
    		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
    	}
    //Assim como somente os índices utilizando o "_"
    for índice, _ := range slice {
    		fmt.Println("Um dos índices do slice é", índice)
    	}
    
    //Outro exemplo
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	slice := []int{20, 21, 22, 23, 1, 13}
    
    	total := 0
    
    	for _, valor := range slice {
    	
    		// mesma coisa que total = total + valor
    		total += valor 
    	
    	}
    	
    	fmt.Println("O valor total é:", total)
    }
    OUTPUT:
    O valor total é: 100
    
    ```

- Fatiando ou deletando de uma fatia

  - Podemos "fatiar" uma slice através da expressão [**low** : **high**], onde _low_ é o índice em que a fatia começa e _high_ é o índice onde ela termina (mas não inclui o próprio índice). Vamos usar um exemplo mais prático, uma pizza:

    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
    
    	fatia := sabores[2:4]
    
    	fmt.Println(fatia)
    }
    OUTPUT:
    [abacaxi quatroqueijos]
    ```

  - Nós podemos ocultar tanto o _low_ quanto o _high_:

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
    
        // Ocultando o low a fatia começa do zero, ou seja, é mesma coisa que sabores[0:4]
    	fatia := sabores[:4]
    
    	fmt.Println(fatia)
    }
    OUTPUT:
    [pepperoni mozzarela abacaxi quatroqueijos]
    ```

  - Quando ocultamos o high a fatia vai até o final.

    ```Go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
    
        // Ocultando o hifh a fatia vai até o fim, ou seja, é mesma coisa que sabores[1:len(sabores)]
    	fatia := sabores[1:]
    
    	fmt.Println(fatia)
    }
    OUTPUT:
    [mozzarela abacaxi quatroqueijos marguerita]
    ```

  - Também podemos ocultar ambos, isso exibe o slice completo

    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
    
        // Ocultando o hifh e o low a fatia começa do zero e vai até o fim, ou seja, é mesma coisa que sabores[0:len(sabores)] ou sabores[0:4]
    	fatia := sabores[:]
    
    	fmt.Println(fatia)
    }
    OUTPUT:
    [pepperoni mozzarela abacaxi quatroqueijos marguerita]
    ```

  - Como deletar um slice? Fatiando!

    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
    	//append adiciona elementos ao final de uma fatia
    	sabores = append(sabores[:2], sabores[4:]...)// os "..." servem para indicar que o que vão ser atribuidos a slice "sabores" é o conteudo de "sabores[4:]" e não a variável. 
    
    	fmt.Println(sabores)
    
    }
    OUTPUT:
    [pepperoni mozzarela marguerita]
    ```

- Função append

  - Faz parte do package builtin
  
  - A função _append_ adiciona elementos no final de uma fatia. Se houver capacidade suficiente no array subjacente, o elemento será inserido após o último elemento e o tamanho será incrementado. Entretanto, se não houver capacidade suficiente, um novo array será criado, todos os elementos existentes serão copiados, o novo elemento será acrescentado no final e a nova fatia será devolvida.
  
    ```GO
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	umaslice := []int{1, 2, 3, 4}
    	outraslice := []int{9, 10, 11, 12}
    
        //Exibe a primeira slice
    	fmt.Println(umaslice)
    
        //Adiciona o números "5, 6, 7, 8"
    	umaslice = append(umaslice, 5, 6, 7, 8)
    
        //Exibe a slice já com os números adicionados
    	fmt.Println(umaslice)
    	
        //Adiciona os "itens" (por isso os "...") da "outraslice" na "umaslice"
    	umaslice = append(umaslice, outraslice...)
    
    	fmt.Println(umaslice)
    }
    OUTPUT:
    [1 2 3 4]
    [1 2 3 4 5 6 7 8]
    [1 2 3 4 5 6 7 8 9 10 11 12]
    ```
