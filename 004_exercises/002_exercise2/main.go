package main

import "fmt"

func main() {
	a := 42
	b := 50

	c := a == b
	fmt.Println(c)

	c = a <= b
	fmt.Println(c)

	c = a >= b
	fmt.Println(c)

	c = a != b
	fmt.Println(c)

	c = a < b
	fmt.Println(c)

	c = a > b
	fmt.Println(c)
}
