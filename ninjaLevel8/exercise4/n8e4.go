package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{11,343,23,124,90,98,56,78,46,21}
	xs := []string{"first","fist","pirst","forst","firlt","fit","ring","word",}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
