package main

import (
	"fmt"
	"log"
)

type customErr struct {
	msg string
}

func main() {
	foo(customErr{msg: "sample error"})
}

func (e customErr) Error() string {
	return fmt.Sprintf("error was encountered: %v", e.msg)
}

// type customErr implements error interface, so we can call foo()
func foo(e error) {
	log.Println(e)
}
