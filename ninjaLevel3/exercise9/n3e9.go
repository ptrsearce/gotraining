package main

import (
	"fmt"
)

func main() {
	favSport := "cricket"
	switch favSport {
	case "cricket":
		fmt.Println("which is ur fav ipl team?")
	case "volleyball":
		fmt.Println("have you watched haikyuu!! anime?")
	case "football":
		fmt.Println("who is ur fav player")
	}
}
