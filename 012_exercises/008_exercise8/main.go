package main

import "fmt"

func main() {
	f := getFunc()
	fmt.Println(f(42))

}

func getFunc() func(x int) int {
	return func(x int) int {
		return x << 2
	}
}
