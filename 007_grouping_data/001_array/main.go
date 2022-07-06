package main

import "fmt"

func main() {
	// Don't use an array if you don't need to.
	// Use a slice instead. Under the hood, a slice references an internal array.
	// Notice that the length of an array is part of its type.
	x := [5]int{1, 2, 3, 4, 5}

	x[3] = 42

	fmt.Println(x)
	fmt.Println(len(x))

	for i, j := range x {
		fmt.Println(i, j)
	}
}
