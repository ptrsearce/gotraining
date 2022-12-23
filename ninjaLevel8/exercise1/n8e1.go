package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "first1",
		Age:   22,
	}

	u2 := user{
		First: "first2",
		Age:   26,
	}

	u3 := user{
		First: "first3",
		Age:   32,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
