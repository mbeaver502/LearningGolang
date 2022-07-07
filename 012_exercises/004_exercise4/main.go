package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   35,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, " and my age is", p.age)
}
