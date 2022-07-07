package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	j := `[
        {
            "First":"James",
            "Last":"Bond",
            "Age":32,
            "Sayings":[
                "Shaken, not stirred",
                "Bond, James Bond"
            ]
    
        }
    ]`

	var xp []person
	err := json.Unmarshal([]byte(j), &xp)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(xp)
}
