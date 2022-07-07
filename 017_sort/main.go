package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type ByAge []person
type ByName []person

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("Unsorted")
	fmt.Println(xi)
	fmt.Println(xs)

	// Standard library's sort package
	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("Sorted")
	fmt.Println(xi)
	fmt.Println(xs)

	p1 := person{
		first: "James",
		age:   32,
	}

	p2 := person{
		first: "Moneypenny",
		age:   27,
	}

	p3 := person{
		first: "Q",
		age:   64,
	}

	p4 := person{
		first: "M",
		age:   56,
	}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByAge(people)) // sort.Sort will use our implementations of Len(), Less(), and Swap()
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

// We can write custom sorting using the sort standard package
//  by implementing the sort.Interface interface
func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Less(i, j int) bool { return p[i].age < p[j].age }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p ByName) Len() int           { return len(p) }
func (p ByName) Less(i, j int) bool { return p[i].first < p[j].first }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
