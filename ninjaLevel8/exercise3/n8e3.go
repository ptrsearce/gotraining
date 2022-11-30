package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "first1",
		Last:  "last1",
		Age:   22,
		Sayings: []string{
			"f1_saying1",
			"f1_saying2",
			"f1_saying3",
		},
	}

	u2 := user{
		First: "first2",
		Last:  "last2",
		Age:   25,
		Sayings: []string{
			"f2_saying1",
			"f2_saying2",
			"f2_saying3",
		},
	}

	u3 := user{
		First: "first3",
		Last:  "last3",
		Age:   35,
		Sayings: []string{
			"f3_saying1",
			"f3_saying2",
			"f3_saying3",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

}
