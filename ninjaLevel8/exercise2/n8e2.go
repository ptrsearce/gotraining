package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s:=`[{"First":"first1","Age":22,"Sayings":["f1_saying1","f1_saying2","f1_saying3"]},{"First":"first2","Age":26,"Sayings":["f2_saying1","f2_saying2","f2_saying3"]},{"First":"first3","Age":32,"Sayings":["f3_saying1","f3_saying2","f3_saying3"]}]`
	fmt.Println(s)

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
