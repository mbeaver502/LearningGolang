package main

import "fmt"

func main() {
	foo(func() string { return "Hello World" })
}

func foo(bar func() string) {
	fmt.Println(bar())
}
