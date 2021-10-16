# Strings (Cadeia de caracteres)

- Strings são sequencias de bytes

- Imutáveis

- Uma string é um "slice of bytes" (fatia de bytes)

- Interpreted string literals (String interpretada)

  - Strings interpretadas, ou seja, podemos usar simbolos para indicar nova linha, tab, negrito, etc.

    ```GO
    package main
    
    import "fmt"
    
    func main() {
    	x := "oi, tudo bem? \n espero \"que\" sim"
    	fmt.Printf(x)
    }
    OUTPUT: 
    oi, tudo bem? 
    espero "que" sim
    ```

- Raw string literals (String "crua")

  - Esse tipo de string não é interpretada, ela "escreve" exatamente o que o usuário passou.

    ```GO
    package main
    
    import "fmt"
    
    func main() {
    	x := `oi, tudo bem? \n espero \"que\" sim`
    	fmt.Printf(x)
    }
    OUTPUT: 
    oi, tudo bem? \n espero \"que\" sim
    ```

    