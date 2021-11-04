# Closure

- Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto.
- Closures nos permitem salvar dados entre function calls e ao mesmo tempo isolar estes dados do resto do código.

 ````Go
 package main
 
 import "fmt"
 
 func main() {
 	a := i()
 
 	fmt.Print(a())
 	fmt.Print(a())
 	fmt.Println(a())
 
 	b := i()
 
 	fmt.Print(b())
 	fmt.Print(b())
 	fmt.Print(b())
 }
 
 func i() func() int {
 	x := 0
 	return func() int {
 		x++
 		return x
 	}
 }
 OUTPUT:
 1
 2
 3
 1
 2
 3
 ````

- Com o closure nós capturamos um scope para usar somente quando quisermos, ou seja, quando a função é criada ela utiliza uma variável que está fora do scope dela, ela utiliza a variável que está no scope da função subjacente. 
- Então cada vez que for feito um closure e capturarmos um scope de fora da função teremos uma cópia diferente desse scope, por isso que o "a" recebe 1,2,3 e "b" também recebe 1,2,3.