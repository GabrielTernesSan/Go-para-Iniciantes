# Enumerando uma slice

- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.

- Exemplos: 

  - Desenrolando uma slice de ints como argumento para a função "soma" anterior.

    ```Go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	si := []int{10, 10, 1, 2, 3, 5}
        
    	// total := soma(si)
        
        /*
        No caso que está acima o preograma não irá funcionar porque o 
        argumento esperado é do tipo int, e não do tipo []int. Já abaixo
        usamos a notação "..." para utilizar os valores que estão dentro da
        slice, que são ints
        */
        
    	total := soma(si...)
    
    	fmt.Println(total)
    }
    
    func soma(x ...int) int {
    
    	soma := 0
    	for _, v := range x {
    		soma += v
    	}
    	return soma
    }
    
    ```

  - Pode-se passar *zero* ou mais valores

    ````Go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	total := soma()
    
    	fmt.Println(total)
    }
    
    func soma(x ...int) int {
    
    	soma := 0
    	for _, v := range x {
    		soma += v
    	}
    	return soma
    }
    ````

  - O parâmetro variádico deve ser o parâmetro final

    ````go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    
    	total, quantos, oi := soma("tarde", 10, 10, 1, 2, 3, 5)
    
    	fmt.Println(total, quantos, oi)
    }
    
    /* 
    Reparem que o elemento variádico aparece primeiro neste caso,
    este programa apresentará o seguinte erro:
    ./prog.go:14:11: syntax error: cannot use ... with non-final parameter x
    */
    func soma(x ...int, s string) (int, int, string) {
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

    