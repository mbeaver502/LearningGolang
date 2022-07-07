package main

import "fmt"

func main() {
	f := foo()

	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //4
}

func foo() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
