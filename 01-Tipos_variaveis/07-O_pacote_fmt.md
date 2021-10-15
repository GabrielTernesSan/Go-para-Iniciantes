# O pacote format (fmt)

- O pacote fmt implementa entrada e saída formatada com funções análogas a printf e scanf do C.

  - Grupo 1: Print 

    - Func Print (a...interface{})(n int, err error)
    - Func Println (a...interface{})(n int, err error) //Adiciona uma linha em branco no final
    - Func Printf (format string, a...interface{})(n int, err error)
      - Format verbs. (%v %T)

  - Grupo 2: Print -> A função Sprint é usada para gerar e retornar uma string formatada

    - Func Sprint(a...interface{}) string

    - Func Sprintf (format string, a...interface{}) string

    - Func Sprint(a...interface{}) string

      ```GO
      package main
      
      import (
      	"fmt"
      )
      
      func main() {
      	quantidade := 10
      	fruta := "banana"
      	texto := fmt.Sprint("Eu tenho ", quantidade, " ", fruta, ".")
      	fmt.Printf("O valor armazenado no texto é: '%v'\n", texto)
      }
      ```

  - Grupo 3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor

    - Func Fprint (w io.Writer, a...interface{})(n int, err error)
    - Func Fprintf (w io.Writer, format string, a...interface{})(n int, err error)
    - Func Fprintln (w io.Writer, a...interface{})(n int, err error)