package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("2 + 4 =", mySum(2, 4))
	fmt.Println("6 + 7 =", mySum(6, 7))
}

func mySum(xi ...int) int {
	sum := 0

	for _, i := range xi {
		sum += i
	}

	return sum
}
