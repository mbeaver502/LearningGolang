package main

import "fmt"

var x int
var y float64
var z uint

func main() {
	//x := 42
	//y := 42.34534
	x = -42
	y = 42.34534
	z = 42

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
