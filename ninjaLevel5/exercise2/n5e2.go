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


	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
