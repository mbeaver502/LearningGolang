package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// go run -race main.go  <-- shows data races

func main() {
	fmt.Println("OS:\t\t\t", runtime.GOOS)
	fmt.Println("Arch:\t\t\t", runtime.GOARCH)
	fmt.Println("Number of CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Number of Goroutines:\t", runtime.NumGoroutine()) // 1 - main

	concurrency()
	raceCondition()
	mutex()
}

func concurrency() {
	// These will run SEQUENTIALLY
	//foo()
	//bar()

	wg.Add(1) // Increment the number of goroutines our WaitGroup will wait for
	go foo()  // Launch a new goroutine and continue execution

	bar() // bar() does not run inside a goroutine, so it's executed immediately in sequence

	wg.Wait() // Wait for all goroutines to be done
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	wg.Done() // Let our WaitGroup know that our goroutine is done
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func raceCondition() {
	counter := 0
	const goroutines = 100

	// We'll encounter a race condition because we're mutating a shared memory space
	// That is, we're attempting to access and modify counter in many goroutines
	// This will lead to unpredictable, undesirable results
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			v := counter

			runtime.Gosched() // Tells scheduler to yield processor to other goroutines

			counter = v + 1

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

func mutex() {
	counter := 0
	const goroutines = 100

	// Can use sync.RWMutex to control reading and/or writing access
	var mu sync.Mutex

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			mu.Lock() // Lock down code access to prevent race condition
			v := counter

			runtime.Gosched() // Tells scheduler to yield processor to other goroutines

			counter = v + 1
			mu.Unlock() // Unlock code access

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter) // Notice that counter is now 100
}
