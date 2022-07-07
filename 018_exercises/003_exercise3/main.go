package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Last  string
	Age   int
}

func main() {
	users := []user{
		{
			First: "James",
			Last:  "Bond",
			Age:   32,
		},
		{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   27,
		},
	}

	fmt.Println(users)

	// Encode the users slice to JSON and then write it out to os.Stdout
	// os.Stdout implements the io.Writer interface
	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("error:", err)
	}
}
