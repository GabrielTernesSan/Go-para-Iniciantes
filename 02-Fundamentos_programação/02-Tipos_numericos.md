# Tipos Numéricos

- Int vc. Float: Números inteiros vs. números com frações.

- Integers

  - Números inteiros

  - int & uint -> "Implementation-specific sizes"

    - Dependem da arquitetura do do computador

    - ```GO
      uint8       o conjunto de todos os inteiros de 8 bits sem sinal (0 to 255)
      uint16      o conjunto de todos os inteiros de 16 bits sem sinal (0 to 65535)
      uint32      o conjunto de todos os inteiros de 32 bits sem sinal (0 to 4294967295)
      uint64      o conjunto de todos os inteiros de 64 bits sem sinal (0 to 18446744073709551615)
      
      int8        o conjunto de todos os inteiros de 8 bits com sinal (-128 to 127)
      int16       o conjunto de todos os inteiros de 16 bits com sinal (-32768 to 32767)
      int32       o conjunto de todos os inteiros de 32 bits com sinal (-2147483648 to 2147483647)
      int64       o conjunto de todos os inteiros de 64 bits com sinal (-9223372036854775808 to 9223372036854775807)
      
      float32     o conjunto de todos os números de ponto flutuante IEEE-754 de 32 bits
      float64     o conjunto de todos os números de ponto flutuante IEEE-754 de 64 bits
      ```

      _REGRA GERAL_: Use sempre "int" e "float64"	

      

  - Todos os tipos numéricos são distintos, exeto: 

    - Byte = uint8
    - rune = int32

  - Tipos são únicos
    - GO é uma linguagem estática
    - Int e int32 não são a mesma coisa
    - Para "misturá-los" é necessário conversão

- Float Point
  - Números racionais ou reais

```GO
package main

import "fmt"

func main() {
	x := 10
	y := 10.0
	fmt.Printf("%v, %T\n%v, %T", x, x, y, y)
}
OUTPUT: 
10
int
10
float64
```



