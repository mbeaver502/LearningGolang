package main

import "fmt"

var x int
var y string
var z bool

func main() {
	// These vars are unassigned, so they're defaulted to their zero values.
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
