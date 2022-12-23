package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "first1",
		last:  "last1",
		favFlavors: []string{
			"f1_fav1",
			"f1_fav2",
			"f1_fav3",
		},
	}

	p2 := person{
		first: "first2",
		last:  "last2",
		favFlavors: []string{
			"f2_fav1",
			"f2_fav2",
			"f2_fav3",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
