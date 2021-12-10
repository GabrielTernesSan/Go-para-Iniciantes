# Afinal o que é bcrypt?

- bcrypt é um método de criptografia do tipo (hash)[https://pt.wikipedia.org/wiki/Fun%C3%A7%C3%A3o_hash] para senhas baseado no Blowfish. Este método
   apresenta uma segurança maior em relação à maioria dos outros métodos criptográficos que é a 
   implementação da variável "custo" que é proporcional à quantidade de processamento necessária 
   para criptografar a senha. O método é conhecido como hash adaptativo às melhorias futuras de 
   hardware por ter esta característica, pois pode permanecer resistente à ataques do tipo 
   "força-bruta" com o tempo usando custos maiores de processamento.

- Ao contrário do método tradicional "crypt" (concebido em 1976), o algoritmo bcrypt não possui as 
  restrições da época que eram pouco processamento e poucos espaço (bytes) para guardar o salt e o
  hash gerado da senha criptografada.

  # O que é hash? 

  - Uma função hash é um algoritmo que mapeia dados de comprimento variável para dados de 
  comprimento fixo. Os valores retornados por uma função hash são chamados valores hash, códigos 
  hash, somas hash (hash sums), checksums ou simplesmente hashes. Um uso é uma estrutura de dados 
  chamada de tabela hash, amplamente usada em software de computador para consulta de dados 
  rápida. Funções hash aceleram consultas a tabelas ou bancos de dados por meio da detecção de 
  registros duplicados em um arquivo grande. Um exemplo é encontrar trechos similares em
  sequências de DNA. Eles também são úteis em criptografia.

  # Como fazer?

- Para fazermos o uso do bcrypt vamos usar duas funções do package do GO (bcrypt)[https://pkg.go.dev/golang.org/x/crypto/bcrypt],
  a GenerateFromPasswordessa função retorna o hash bcrypt da senha com o custo computacional fornecido, e se o custo computacional 
  fornecido for menor que MinCost (MinCost int = 4), o custo será definido como DefaultCost. A segunda função é a 
  CompareHashAndPassword que compara uma senha com hash bcrypt com seu possível equivalente em texto simples. Retorna nulo em caso
  de sucesso ou erro em caso de falha.

```
    package main

    import "fmt"
    import "golang.org/x/crypto/bcrypt"

    func main() {
        senha := "GolangEhMuitoDaora"
        senhaerrada := "GolangehMuitoDaora"

        sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(string(sb))

        fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
        fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

    }
```

    10 = (custo computacional padrão)[https://cs.opensource.google/go/x/crypto/+/5770296d:bcrypt/bcrypt.go]
    `sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)`