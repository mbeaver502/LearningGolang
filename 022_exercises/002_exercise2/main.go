package main

import "fmt"

func main() {
	// Original declaration
	// cs := make(chan<- int)

	// Make it bidirectional
	cs := make(chan int)
	go func() { cs <- 42 }()

	// We cannot read from a send-only channel
	fmt.Println(<-cs)

	fmt.Println("----")
	fmt.Printf("%T\n", cs)
}
