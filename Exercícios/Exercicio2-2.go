package main

import (
	"fmt"
)

func main() {
	a := (29 == 9) // false
	b := (29 != 9) // true
	c := (29 <= 9) // false
	d := (29 < 9)  // false
	e := (29 >= 9) // true
	f := (29 > 9)  // true
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
