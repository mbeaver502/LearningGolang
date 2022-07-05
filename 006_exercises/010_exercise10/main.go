package main

import "fmt"

func main() {
	fmt.Println("true && true", (true && true))   // true
	fmt.Println("true && false", (true && false)) // false
	fmt.Println("true || true", (true || true))   // true
	fmt.Println("true || false", (true || false)) // true
	fmt.Println("!true", (!true))                 // false
}
