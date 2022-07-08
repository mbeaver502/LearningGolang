package main

import (
	"fmt"

	"github.com/mbeaver502/LearningGolang/026_exercises/001_exercise1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido.age)
}
