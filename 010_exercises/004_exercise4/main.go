package main

import "fmt"

func main() {
	s := struct {
		first, last string
		age         int
	}{
		first: "Jack",
		last:  "Bauer",
		age:   35,
	}

	fmt.Println(s)
}
