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

	bs, err := toJSON(p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func toJSON(a any) ([]byte, error) {
	bs, err := json.Marshal(a)

	//err = errors.New("force an error for example purposes")

	if err != nil {
		return []byte{}, fmt.Errorf("toJSON: there was an error while marshalling: %v", err)
	}

	return bs, nil
}
