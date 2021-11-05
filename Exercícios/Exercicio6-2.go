package main

import "fmt"

func main() {
	total := soma(1, 2, 3, 4)

	fmt.Println(total)

	num := []int{1, 2, 3, 4, 5, 6}
	total2 := soma2(num...)

	fmt.Println(total2)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func soma2(x ...int) int {
	soma2 := 0
	for _, valor := range x {
		soma2 += valor
	}
	return soma2
}
