package main

import "fmt"

//const a = 42
//const b = 42.78
//const c = "James Bond"

// Constants can be untyped or explicitly typed
const (
	a int = 42
	b     = 42.78
	c     = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
