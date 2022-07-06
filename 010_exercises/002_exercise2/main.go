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

	m := make(map[string]person)

	m[p1.last] = p1
	m[p2.last] = p2

	fmt.Println(m)

	for _, val := range m {
		fmt.Println(val.first, val.last)
		for _, f := range val.flavors {
			fmt.Println(f)
		}
	}
}
