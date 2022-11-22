package main

import (
	"fmt"
)

func main() {
	arr := [5]int{11,12, 13, 14, 15}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
