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

