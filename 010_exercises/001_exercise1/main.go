package main

import "fmt"

type person struct {
	first, last string
	flavors     []string
}

func main() {
	p1 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"Chocolate", "Hazelnut"},
	}

	p2 := person{
		first:   "Miss",
		last:    "Moneypenny",
		flavors: []string{"Strawberry", "Rocky Road"},
	}

	fmt.Println(p1.first, p1.last)
	for _, f := range p1.flavors {
		fmt.Println(f)
	}

	fmt.Println(p2.first, p2.last)
	for _, f := range p2.flavors {
		fmt.Println(f)
	}
}
