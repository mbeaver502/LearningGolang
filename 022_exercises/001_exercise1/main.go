package main

import "fmt"

func main() {
	// deadlock()
	lambda()
	buffer()
}

func deadlock() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

func lambda() {
	c := make(chan int)

	go func() { c <- 42 }()

	fmt.Println(<-c)
}

func buffer() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
