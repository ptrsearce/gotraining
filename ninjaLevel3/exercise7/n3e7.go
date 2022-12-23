package main

import (
	"fmt"
)

func main() {
	x := "really?"

	if x == "really?" {
		fmt.Println(x)
	} else if x == "not so" {
		fmt.Println("hushh..", x)
	} else {
		fmt.Println("hmmm..")
	}
}
