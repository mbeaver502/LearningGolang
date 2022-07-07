package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		foo()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		bar()
		wg.Done()
	}()

	wg.Wait()
}

func foo() {
	fmt.Println("Greetings from foo")
}

func bar() {
	fmt.Println("Greetings from bar")
}
