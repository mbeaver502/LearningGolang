package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak(s string)
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
	}

	p.speak("Shaken, not stirred")

	// cannot use p directly because there is no func (p person) speak(s string)
	// the only method available is for a *person type, not a person type
	// in other words, person does *not* implement the human interface
	// however, *person *does* implement the human interface
	saySomething(&p, "Shaken, not stirred")
}

func (p *person) speak(s string) {
	fmt.Println(p.first, p.last, "says", s)
}

func saySomething(h human, s string) {
	h.speak(s)
}
