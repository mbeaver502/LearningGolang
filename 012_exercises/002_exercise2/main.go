package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}

	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(i ...int) int {
	sum := 0

	for _, n := range i {
		sum += n
	}

	return sum
}

func bar(i []int) int {
	return foo(i...)
}
