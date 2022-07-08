package main

import "fmt"

func main() {
	c := make(chan int)

	// Load some message onto our channel
	go func() { c <- 42 }()

	// Try to take some message off our channel
	v, ok := <-c
	fmt.Println(v, ok) // 42 true

	// We're done with our channel, so close it to further messages
	close(c)

	// Try to take some message off our channel, which will fail
	v, ok = <-c
	fmt.Println(v, ok) // 0 false (0 is the zero value for int)
}
