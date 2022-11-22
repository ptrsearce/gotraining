package main

import (
	"fmt"
)

func main() {
	bday := 2000
	for {
		if bday > 2022 {
			break
		}
		fmt.Println(bday)
		bday++
	}
}
