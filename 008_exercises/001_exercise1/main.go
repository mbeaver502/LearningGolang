package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
