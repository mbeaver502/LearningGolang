package main

import (
	"fmt"
)

// Channels are all about sharing memory by communicating
// Bidrectional channels can both send and receive
// Directional channels can only send or receive
func main() {
	//channelsBlockDoesNotWork()
	channelsBlockGoroutine()
	channelsBlockBuffer()

	directionalChannel()
	channelsWithFuncs()

	rangeOnChannel()
	selectChannel()
}

// This will not run because channels block
// Sending and receiving on a channel is a blocking action
// Send will block until receive is ready, and vice versa
func channelsBlockDoesNotWork() {
	// Create a channel that can communicate integers
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

// This will work because now there's two goroutines -- one to send and one to receive
func channelsBlockGoroutine() {
	// Create a channel that can communicate integers
	c := make(chan int)

	// Launch goroutine to send a message onto our channel
	go func() { c <- 42 }()

	// Now we can listen for the message on our channel
	fmt.Println(<-c)
}

// We can use buffered channels to temporarily store messages on the channel
// In general, stay away from buffered channels
func channelsBlockBuffer() {
	// Create a channel that can communicate integers
	// This time our channel is buffered
	// In this case, c can buffer one message
	c := make(chan int, 1)

	c <- 42
	// c <- 43 // This buffer would be unsuccessful since our channel buffer is full

	// Now we can listen for the message on our channel
	fmt.Println(<-c)
}

// Channels can be directional
func directionalChannel() {
	b := make(chan int)      // bidirectional chan
	s := make(chan<- int, 2) // send-only chan
	r := make(<-chan int, 2) // receive-only chan

	fmt.Printf("%T\n", b) // chan int
	fmt.Printf("%T\n", s) // chan<- int
	fmt.Printf("%T\n", r) // <-chan int

	// Assignment works if we're going from general to specific (chan int => <-chan int)
	r = b
}

func channelsWithFuncs() {
	c := make(chan int)
	fmt.Println("starting")

	go sender(c)
	receiver(c) // this is a blocking call, and we'll block until we get a message over c

	fmt.Println("exiting")
}

// We can have channels as function parameters
func sender(c chan<- int) {
	c <- 50
}

func receiver(c <-chan int) {
	fmt.Println(<-c)
}

func rangeOnChannel() {
	c := make(chan int)

	fmt.Println("starting")

	// We can send a bunch a values over a channel and then close it when we're done
	go func(x chan<- int) {
		for i := 0; i < 100; i++ {
			x <- i
		}

		// Close the channel -- this closes the c declared in rangeOnChannel()
		// If we don't close, we'll deadlock because there's someone listening for messages
		close(x)
	}(c)

	// We can range over a channel until it is closed
	func(y <-chan int) {
		for v := range y {
			fmt.Println(v)
		}
	}(c)

	fmt.Println("exiting")
}

func selectChannel() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	fmt.Println("starting")

	// send
	go func(e, o chan<- int, q chan<- bool) {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				e <- i
			} else {
				o <- i
			}
		}

		close(q)
	}(even, odd, quit)

	// receive
	func(e, o <-chan int, q <-chan bool) {
		for {
			// We can use select to pull messages off multiple channels
			select {
			case v := <-e:
				fmt.Println("even:", v)
			case v := <-o:
				fmt.Println("odd:", v)

			// ok will be false if no value was received
			case v, ok := <-q:
				if !ok {
					fmt.Println("quit:", v, ok)
				} else {
					fmt.Println("quit:", v, ok)
				}

				return
			}
		}
	}(even, odd, quit)

	fmt.Println("exiting")
}
