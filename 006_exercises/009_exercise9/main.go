package main

import "fmt"

func main() {
	favSport := "soccer"

	switch favSport {
	case "soccer":
		fmt.Println("You like soccer.")
	case "football":
		fmt.Println("You like football.")
	default:
		fmt.Println("You like something else.")
	}
}
