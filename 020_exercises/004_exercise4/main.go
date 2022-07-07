package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)

		// solve race condition with mutex lock
		go func() {
			mu.Lock()
			v := counter

			runtime.Gosched()

			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
