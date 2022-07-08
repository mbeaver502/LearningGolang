package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("exiting")
}

func gen() <-chan int {
	c := make(chan int)

	// This gets launched as a goroutine
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		// Close c after we load messages onto it so we avoid deadlock
		close(c)
	}()

	// c gets returned after we launch our anonymous goroutine
	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
