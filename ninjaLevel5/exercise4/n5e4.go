package main

import (
	"fmt"
)

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "first1",
		friends: map[string]int{
			"first2": 555,
			"first3": 777,
			"first4": 888,
		},
		favDrinks: []string{
			"fav1",
			"fav2",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
