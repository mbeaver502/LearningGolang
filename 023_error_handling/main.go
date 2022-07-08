package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type myError struct {
	msg string
	val int
}

// In Go, we don't use try-catch-finally
// In Go, we can return errors from funcs, which makes catching errors convenient and easy
// There are ways to handle *exceptional* errors, but most "common" errors are handled simply
// Go has a built-in error interface that exposes the func Error() string
// Any type that implements Error() string implicitly implements the error interface
// We also have access to the standard library's errors package
// Always, always, always check errors (with minor exception)
func main() {
	//err := errors.New("my error")
	//fmt.Println(err.Error())
	//fmt.Println(err)

	//exPrint()
	//exFile()
	//exLog()

	//exRecover()
	//fmt.Println("Execution returned to main")

	// The return value of fmt.Errorf is of type error
	// Under the hood, fmt.Errorf calls errors.New()
	e := fmt.Errorf("we can use fmt.Errorf to format error strings %v", 42)
	fmt.Println(e)

	err := myError{
		msg: "my error message",
		val: 42,
	}

	fmt.Println(err)
}

func exPrint() {
	// By convention, don't bother checking errors on fmt.Println()
	if n, err := fmt.Println("Hello World"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("bytes:", n)
	}
}

func exFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}

func exLog() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("Error happened:", err) // simply puts the error on os.Stdout
		log.Println("Error happened:", err) // same as fmt.Println(), except we get a timestamp
		//log.Fatalln(err)                    // log.Fataln() logs the error and does os.Exit(1) (no defers are run)
		//log.Panicln(err) // same as fmt.Println() followed by panic() (defers are run)
		//panic(err) // stops normal execution and runs deferred funcs (panic can be caught by recover)
	}

	defer f2.Close()

	fmt.Println("Check log.txt")
}

// PANIC WILL STOP NORMAL EXECUTION AND START UNWINDING THE DEFERRED FUNCTIONS CALL LIST
// RECOVER IS ONLY USEFUL INSIDE DEFERRED FUNCTIONS
func exRecover() {
	defer func() {
		// r holds the value passed into panic()
		if r := recover(); r != nil {
			fmt.Println("Recovered!", r)
		}
	}()

	exPanic(0)

	fmt.Println("Returning normally from exRecover")
}

func exPanic(i int) {
	if i > 3 {
		fmt.Println("Panic!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in exPanic", i)

	fmt.Println("Printing in exPanic", i)
	exPanic(i + 1)

	fmt.Println("Returning normally from exPanic") // will not execute if we panic
}

func (e myError) Error() string {
	return fmt.Sprintf("%v: %d", e.msg, e.val)
}
