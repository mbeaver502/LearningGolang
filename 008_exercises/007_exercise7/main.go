package main

import "fmt"

func main() {
	ss := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hellooooo, James"},
	}

	fmt.Println(ss)

	for _, s := range ss {
		fmt.Println(s)

		for _, t := range s {
			fmt.Println(t)
		}
	}
}
