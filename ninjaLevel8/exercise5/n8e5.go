package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

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
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("-------")

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("-------")

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

}
