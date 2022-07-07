package main

import (
	"encoding/json"
	"fmt"
)

// JSON-to-Go website will convert JSON-encoded strings to Go structs
// https://mholt.github.io/json-to-go/

// In order for our fields to be encoded, they need to be exported (capitalized)
// We can control the field name in the JSON encoding by using `json:"name"` tags
type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"Whatever"`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   35,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	// This will give back empty JSON objects if the struct fields are unexported
	// Struct fields are encoded by json.Marshal() if the fields are exported, hence the capitals
	j, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(j))

	// We can also unmarshal JSON-encoded strings
	// Note that the second param is a pointer to the destination data structure
	var xp []person
	err = json.Unmarshal(j, &xp)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(xp)
}
