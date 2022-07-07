package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		// Solve race condition using sync/atomic -- but prefer mutex locking when possible
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
