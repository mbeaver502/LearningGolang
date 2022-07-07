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

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	// alternatively: (*p).first, (*p).last, (*p).age
	p.first = "Miss"
	p.last = "Moneypenny"
	p.age = 26
}
