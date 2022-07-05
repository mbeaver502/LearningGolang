package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	a := 1
	b := 10

	// We can run a for-loop for a simple condition, like a while-loop
	for a < b {
		fmt.Println(a, b)
		a *= 2
	}

	s := "Hello World!"

	// We can run a for-loop over a range
	// We can use _ to discard the index or the ranged value
	for i, c := range s {
		fmt.Println(i, c, string(c))
	}

	x := 1

	// We can do an infinite loop
	for {
		if x > 10 {
			break // We can break out of a for-loop early
		} else if x%2 == 0 {
			x++
			continue // We can use continue to resume a loop early
		}

		fmt.Println(x)
		x++
	}
}
