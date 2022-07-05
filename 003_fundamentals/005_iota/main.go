package main

import "fmt"

// We can use iota to automatically increment starting at 0 (or 0 + some value)
const (
	a = iota
	b
	c
)

// iota resets at the next const
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
