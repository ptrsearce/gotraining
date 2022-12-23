package main

import (
	"fmt"
)

type mywish int

var x mywish

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
