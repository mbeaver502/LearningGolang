package main

import "fmt"

func main() {
	a := 5

	if a < 3 {
		fmt.Println("Less than 3")
	} else if a == 3 {
		fmt.Println("Equals 3")
	} else {
		fmt.Println("Greater than 3")
	}
}
