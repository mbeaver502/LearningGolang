package main

import "fmt"

// Type "person" has two values/fields, in this case both of type "string"
type person struct {
	first, last string
}

// We can embed structs inside other structs
type agent struct {
	// We embed the person struct inside the agent struct -- person is an "embedded field"
	// In this case, the unqualified type name (person) also acts as the field name
	person
	licenseToKill bool
}

func main() {
	// Create a variable "jack" that holdes values of type "agent"
	jack := agent{
		// If we want to initialize person's fields, we need to be explicit that we're accessing person
		person: person{
			first: "Jack",
			last:  "Bauer",
		},
		licenseToKill: true,
	}

	// We can use anonymous structs for one-off purposes
	p := struct {
		first, last string
		age         int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(jack)

	// Notice how person.first and person.last are promoted to the enclosing agent type.
	fmt.Println(jack.first, jack.last, jack.licenseToKill)

	fmt.Println(p)

	fmt.Printf("%T\n%T\n", jack, p)
}
