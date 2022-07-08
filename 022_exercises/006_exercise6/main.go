package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)
	receive(c)

	fmt.Println("exiting")
}

// Load a bunch of messages onto our channel
// Close the channel when we're done so our receiver doesn't deadlock
func send(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}

// Keep taking messages of our channel until the channel is closed
// Loop will deadlock if the channel is not closed and there are no messages
func receive(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
