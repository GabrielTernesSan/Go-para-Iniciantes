# Operador curto de declaração

- := é chamado de Gopher por parecer uma marmota

- Uso

  - Tipagem automática;

  - Só pode repetir se houverem variáveis novas:

    ```GO
    package main
    
    import "fmt"
    
    func main() {
        x := 10
        fmt.Println("x: %v", x)
        //x := 20 dará erro, mas se houver uma variável nova sendo declarada..
        x, z := 20, 30 // "z" variável nova
    }
    ```

  - Só funciona dentro do escopo, não possui içamento como em algumas linguagens como JavaScript.

    ```Go
    package main
    
    import "fmt"
    
    func main() {
        x := 10
        fmt.Println("Valor de x: %v, Tipo: %T", x, x)
        fmt.Println("Valor de z: %v, Tipo: %T", z, z) //fora de escopo. A variável "z" não foi declarada (ela ainda não existe).
        x, z := 20, 30 // Nesta linha a variável "z" foi declarada
        fmt.Println("x: %v, %T", x, x)
        fmt.Println("z: %v, %T", z, z)
    }
    ```

    