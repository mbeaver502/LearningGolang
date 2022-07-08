package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Never say never",
		},
	}

	bs, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("JSON error:", err)
	}

	fmt.Println(string(bs))
}
