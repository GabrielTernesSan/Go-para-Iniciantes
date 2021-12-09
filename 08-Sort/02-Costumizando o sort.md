
  # Costomizando o sort
- Como fazer o seu próprio sort
    Para fazermos nosso próprio sort vamos usar a func Sort do package sort. Vamos precisar de um [sort.Interface](https://pkg.go.dev/sort#Interface) `type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }`. Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como nos exemplos anteriores.

```
    package main

    import (
        "fmt"
    )

    type carro struct {
        nome string
        potencia int
        consumo int
    }

    type ordenarPorPotencia []carro
    
    //Len é o número de elementos da coleção.
    func (x ordenarPorPotencia) Len() int {return len(x)}

    // Less relata se o elemento com índice i deve ser classificado antes do elemento com índice j.
    //Se ambos Less (i, j) e Less (j, i) são falsos, então os elementos no índice i e j são considerados iguais. O sort pode
    //colocar elementos iguais em qualquer ordem no resultado final, enquanto Stable preserva a ordem de entrada original de
    //elementos iguais.
    func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }

    //Swap troca os elementos com os índices i e j.
    func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

    type ordenarPorConsumo []carro

    func (x ordenarPorConsumo) Len() int           { return len(x) }
    func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
    func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

    // Quanto maior for o consumo, mais lucro o dono do posto vai ter.
    type ordenarPorLucroProDonoDoPosto []carro
    func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
    func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
    func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

    func main() {

	carros := []carro{carro{"Chevete", 50, 8},
		carro{"Porsche", 300, 5},
		carro{"Fusca", 20, 30},
	}

	fmt.Println("Inicial:\n", carros)
    //[{Chevete 50 8} {Porsche 300 5} {Fusca 20 30}]

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potência:\n", carros)
    //[{Fusca 20 30} {Chevete 50 8} {Porsche 300 5}]

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Consumo:\n", carros)
    //[{Fusca 20 30} {Chevete 50 8} {Porsche 300 5}]

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	fmt.Println("Lucro:\n", carros)
    //[{Porsche 300 5} {Chevete 50 8} {Fusca 20 30}]
    }
```