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
type ByLastName []user
type ByFirstName []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (n ByLastName) Len() int           { return len(n) }
func (n ByLastName) Less(i, j int) bool { return n[i].Last < n[j].Last }
func (n ByLastName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (n ByFirstName) Len() int           { return len(n) }
func (n ByFirstName) Less(i, j int) bool { return n[i].First < n[j].First }
func (n ByFirstName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func main() {
	users := []user{
		{
			First: "James",
			Last:  "Bond",
			Age:   32,
			Sayings: []string{
				"Shaken, not stirred",
				"Bond, James Bond",
			},
		},
		{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   27,
			Sayings: []string{
				"I don't know???",
			},
		},
		{
			First: "Doctor",
			Last:  "No",
			Age:   53,
			Sayings: []string{
				"I'll get you, Bond",
				"Does he actually say that???",
			},
		},
		{
			First: "Q",
			Last:  "Classified",
			Age:   60,
			Sayings: []string{
				"I really like gadgets.",
			},
		},
	}

	sort.Sort(ByAge(users))
	print(&users)

	fmt.Println("----------")
	fmt.Println()

	sort.Sort(ByLastName(users))
	print(&users)

	fmt.Println("----------")
	fmt.Println()

	sort.Sort(ByFirstName(users))
	print(&users)
}

func print(users *[]user) {
	for _, u := range *users {
		fmt.Println(u.First, u.Last, u.Age)

		sayings := u.Sayings
		sort.Strings(sayings)

		for _, saying := range sayings {
			fmt.Println("\t", saying)
		}

		fmt.Println()
	}
}
