package main

import "fmt"

func main() {
	f := func(x int) int { return x << 2 }

	fmt.Println(f(42))
}
