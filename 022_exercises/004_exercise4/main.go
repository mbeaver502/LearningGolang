package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("exiting")
}

func gen(q chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		q <- 1
		close(c)
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Value from chan c:", v)
		case v := <-q:
			fmt.Println("Value from chan q:", v)
			return
		}
	}
}
