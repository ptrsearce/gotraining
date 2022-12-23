package main

import (
	"fmt"
)

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(s1)
	fmt.Println(s2)

	s := [][]string{s1, s2}
	fmt.Println(s)

	for i, v := range s {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf(" %v \t value: \t %v \n", j, val)
		}
	}
}
