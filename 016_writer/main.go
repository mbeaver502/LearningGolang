package main

import (
	"fmt"
	"io"
)

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Fprintln(&p, p)
	io.WriteString(&p, "hello world!")
}

// The person type now implements the io.Writer interface
// Since the person type implements io.Writer, we can pass a person-pointer to fmt.Fprintln()
func (p *person) Write(b []byte) (n int, err error) {
	return fmt.Println(p.first, p.last, "from inside Write. Argument:", string(b))
}
