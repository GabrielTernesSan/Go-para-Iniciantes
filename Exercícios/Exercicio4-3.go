package main

import "fmt"

func main() {
	slice := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	fmt.Println("Do primeiro ao terceiro item do slice", slice[:4])
	fmt.Println("Do quinto ao último item do slice", slice[4:])
	fmt.Println("Do segundo ao sétimo item do slice", slice[1:7])
	fmt.Println("Do terceiro ao penúltimo item do slice", slice[3:9])
	fmt.Println("Do terceiro ao penúltimo item do slice", slice[3:(len(slice)-1)])
}
