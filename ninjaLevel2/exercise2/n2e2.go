package main

import (
	"fmt"
)

func main() {
	a := (7 == 7)
	b := (7 <= 11)
	c := (7 >= 11)
	d := (7 != 11)
	e := (7 < 11)
	f := (7 > 11)
	fmt.Println(a, b, c, d, e, f)
}
