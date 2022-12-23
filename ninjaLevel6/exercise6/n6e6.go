package main

import (
	"fmt"
)

func main() {

	func() {
		for i := 0; i < 25; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}
