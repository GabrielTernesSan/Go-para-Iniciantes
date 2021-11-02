# Funções

- Sintaxe

  - Qual a utilidade de funções? 

    - Abstrair funções: 

      ````GO
      package main
      
      import (
      	"fmt"
      )
      
      func main(){
          s := "abc"
          
          fmt.Println(len(s))
      }
      ````

      ​	Tanto a função "len" como a função "Println", não foram desenvolvidas neste programa elas estão abstraídas e foram importadas de algum lugar. Neste código não precisamos saber como a string é contada ou como ela é exibida na tela.

    - Reutilização de código

      ​	Imagine que toda fez que você precise contar quantos elementos a string possui precise escrever 10-15 linhas de código, ou exibir na tela algum valor e ter que escrever 20-30 linhas. Com a utilização de funções estas linhas de código podem ser reutilizadas apenas passando um parâmetro.

    - Toda função segue esta estrutura

      - func nome(parâmetro) (retorno) { Código}

    - A diferença entre parâmetros e argumentos:

      - Funções são definidas com parâmetros, ou seja, quando criamos uma função defino um parâmetro.
      - Funções são chamadas com argumentos, quando eu chamo uma função eu passo um argumento.

    - Parâmetro pode ser ...variádico.

    - Exemplos:

      - Função básica

        ````Go
        package main
        
        import (
        	"fmt"
        )
        
        func main() {
        	// Toda vez que a função é chamada ela exibe "Oi, bom dia!"
        	basica()
        }
        
        func basica() {
        	fmt.Println("Oi, bom dia!")
        }
        
        ````

      - Função que aceita um argumento.

        ````Go
        package main
        
        import (
        	"fmt"
        )
        
        func main() {
            /*Quando o argumento for igual a "manhã" será exibido
            "Oi, bom dia!"*/
        	argumento("manhã")
            /*Quando o argumento for igual a "tarde" será exibido
            "Oi, boa tarde!"*/
        	argumento("tarde")
            /*Quando o argumento for diferente das opções acima será exibido
            "Oi, boa noite!"*/
        	argumento("jfdhjf")
        }
        
        func basica() {
        	fmt.Println("Oi, bom dia!")
        }
        
        func argumento(s string) {
        	if s == "manhã" {
        		fmt.Println("Oi, bom dia!")
        	} else if s == "tarde" {
        		fmt.Println("Oi, boa tarde!")
        	} else {
        		fmt.Println("Oi, boa noite!")
        	}
        }
        ````

      - Função com retorno.

        ```go
        package main
        
        import (
        	"fmt"
        )
        
        func main() {
        // Argumento
        	valor := soma(10, 10)
        // Exibe o resultado
        	fmt.Println(valor)
        }
        // Retorna a soma dos dois números passados como argumento
        func soma(x, y int) int {
        	return x + y
        }
        ```

      - Função com múltiplos retornos e parâmetro variádico.

        ````go
        package main
        
        import (
        	"fmt"
        )
        
        func main() {
        	//total = a soma dos elementos
            //quantos = len(x) que exibirá a quantidade de elementos
        	total, quantos, oi := soma(10, 10, 1, 2, 3, 5)
        
        	fmt.Println(total, quantos, oi)
        }
        //Nome(Parâmetro Variádico -> slice de ints) (Retorno)
        func soma(x ...int) (int, int, string) {
        	soma := 0
            //x é uma slice, por isso, podemos usar 'range'
        	for _, v := range x {
        		soma += v
        	}
        	return soma, len(x), "Bom dia!"
        }
        ````

        Quando temos uma função com parâmetros variádicos o elemento variádico dp parâmetro tem que ser o último, necessariamente deve ser o último, você pode ter quantos elementos quiser antes do elemento variádico. Mais um exemplo:

        ````GO
        package main
        
        import (
        	"fmt"
        )
        
        func main() {
        
        	total, quantos, oi := soma("tarde", 10, 10, 1, 2, 3, 5)
        
        	fmt.Println(total, quantos, oi)
        }
        //	nome(1º elemento, elemento variádico) (Retorno)
        func soma(s string, x ...int) (int, int, string) {
        	oi := ""
        	if s == "manhã" {
        		oi = "Oi, bom dia!"
        	} else if s == "tarde" {
        		oi = "Oi, boa tarde!"
        	} else {
        		oi = "Oi, boa noite!"
        	}
        	soma := 0
        	for _, v := range x {
        		soma += v
        	}
        	return soma, len(x), oi
        }
        
        ````

        