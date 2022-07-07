package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 3, 6, 8, 1, 2, 0, 1, 3, 4, 6}
	xs := []string{"bob", "john", "alice", "steve", "alan", "billy"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
