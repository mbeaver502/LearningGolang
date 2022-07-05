package main

import "fmt"

// Variable x is of type bool.
var x bool

func main() {
	// The zero value of bool is false.
	fmt.Println(x)

	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a, "==", b, "?", a == b)
	fmt.Println(a, "<", b, "?", a < b)
	fmt.Println(a, ">", b, "?", a > b)
}
