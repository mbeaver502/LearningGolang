package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 6, "hi"
}
