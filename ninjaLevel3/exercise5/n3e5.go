package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("The modulus of %v when divided by 4 is %v\n", i, i%4)
	}
}
