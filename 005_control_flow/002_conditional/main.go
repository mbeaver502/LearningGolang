package main

import "fmt"

func main() {
	if true {
		fmt.Println("This gets printed.")
	}

	if false {
		fmt.Println("This does NOT get printed.")
	}

	a := 2
	b := 2

	if a == b {
		fmt.Println("This gets printed, too.")
	}

	// We can initialize a variable that will have local scope to the if-condition's block.
	if x := 42; x > 40 {
		fmt.Println("We can initialize a variable in an if-condition.")
		fmt.Println("x is limited to the scope of this if-condition's block.", x)
	}

	a = 10
	b = 20

	// We can have multiple branches.
	// We can have as many else if statements as we like.
	if a > b {
		fmt.Println("do something")
	} else if a == b {
		fmt.Println("do something else")
	} else {
		fmt.Println("do the default")
	}

	a = 42

	// We can use a switch statement with pre-defined cases.
	switch a {
	case 37, 38, 39:
		fmt.Println("We can have multiple cases in one.")
	case 40:
		fmt.Println("a is 40")
	case 41:
		fmt.Println("a is 41")
	case 42:
		fmt.Println("a is 42")
		//fallthrough // We can use fallthrough to slide down to the next case
	default:
		fmt.Println("a is not 40, 41, or 42")
	}
}
