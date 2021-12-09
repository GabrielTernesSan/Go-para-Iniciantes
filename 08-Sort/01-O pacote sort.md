# O pacote sort

- Sort serve para ordenar slices
- Sort altera o valor original!
- [Documentação](https://pkg.go.dev/sort) : "O package sort fornece elementos primitivos para ordenar slices e coleções definidas pelo usuário."
- Como faz?
  - Strings
  
```
    package main
    import (
	"fmt"
	"sort"
    )

    func main() {

	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println(ss)
    //[abóbora maçã laranja beringela berinjela]
	
	sort.Strings(ss)
	
	fmt.Println(ss)
	//[abóbora beringela berinjela laranja maçã]
    }
```
  - Ints

```
    package main

    import (
        "fmt"
        "sort"
    )

    func main() {

        si := []int{123, 987, 324, 876, 234, 987, 234, 76}

        fmt.Println(si)
        //[123 987 324 876 234 987 234 76]

        sort.Ints(si)

        fmt.Println(si)
        //[76 123 234 234 324 876 987 987]

    }
```