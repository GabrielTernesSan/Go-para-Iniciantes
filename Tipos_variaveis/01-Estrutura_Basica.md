# Terminologia

- **Packages**: são coleções de funções pré-prontas (ou não) que você pode utilizar. Em algumas linguagens são chamados de bibliotecas.

  - Exemplo:

    ```GO
    package main
    ```

- **Imports**: O import "chama" uma função de um package. A função "println" não foi criada, ela foi usada através da importação do package format.

  - Exemplo: 

    ```Go
    import "fmt"
    	func main() {
    		fmt.Println("Olá Mundo")
    }
    ```


- **Variáveis**: Uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão.

- **Expressões**: Qualquer coisa que produz um resultado.

- **Declaração**: Uma linha de código, uma instrução que forma uma ação, formada de expressões.

- **Inicialização**: primeiro valor atribuído a uma variável

- **Atribuição**: alteração.

  ```Go
  import "fmt"
  	func main() {
          //Declaração de variáveis
          var int a 
          var int b
          var int c
          //Expressão
          c = a + b
          
  		fmt.Println(c)
  }
  ```

- **Keywords** (Palavras-chave): São termos reservados que não podem ser usadas como identificadores.
  - Algumas palavras reservadas:

```GO
break  default  func  interface  select
case   defer	go	  map		 struct
chan   else     goto  package    switch
const  if       range type       continue
for    import   return           var
```

- **Valor zero**: O que é valor zero?

  -  Valor zero é o valor que se encontra em uma variável antes dela ser _inicializada_ pelo usuário.

  - Os zeros:

    ````GO
    Int: 0         Boolean: false
    Float: 0.0     string: ""
    Ponteiros, functions, slices, maps: nil
    ````

    
