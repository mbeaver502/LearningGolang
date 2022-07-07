package main

import "fmt"

type person struct {
	first, last string
}

type agent struct {
	person
	ltk bool
}

// Any type that implements these functions can also be considered to be of type human
// We don't necessarily care *how* these functions are implemented
// We care that we are *guaranteed* that these functions are available
type human interface {
	speak()
}

func main() {
	defer bar("deferred funcs execute at end of func body, on return, or on panic")
	defer bar("Notice that deferred funcs execute in LIFO order")
	defer bar("defer is handy for guaranteeing that open resources are eventually closed")
	defer fmt.Println(sumInts(6, 7, 8, 9, 0))

	foo()
	bar("Hello from bar!")
	fmt.Println(woo("Shenanigans"))
	fmt.Println(revCat("James", "Bond"))

	fmt.Println(sumInts())              // Value will be nil under the hood
	fmt.Println(sumInts(1))             // A new slice is created under the hood
	fmt.Println(sumInts(1, 2, 3, 4, 5)) // A new slice is created under the hood

	// We can use ... to unfurl a slice so that we can use it for a variadic parameter
	// xi is already []int, so no new slice is created under the hood
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(sumInts(xi...))

	p1 := agent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := agent{
		person: person{
			first: "Miss",
			last:  "Monneypenny",
		},
		ltk: false,
	}

	p3 := person{
		first: "Doctor",
		last:  "No",
	}

	xh := []human{p1, p2, p3}

	for _, h := range xh {
		doHumanStuff(h)
	}

	p1.kill()
	p2.kill()

	// We can also use anonymous functions
	// Anonymous functions can also take in parameters and arguments
	func(x, y int) {
		fmt.Println("Hello from my anonymous func!", x, y)
	}(2, 5)

	// Functions are types, so we can have functions as expressions
	fe := func(x, y int) int { return x << y }
	fmt.Println(fe(5, 3))

	// We can also return functions from other functions
	rf := funcReturner()
	fmt.Println(rf(10, 3))
	fmt.Println(funcReturner()(5, 3)) // No need to assign the func to a var
	fmt.Printf("%T\n", rf)

	// We can also pass functions as arguments to other functions (i.e., as a callback)
	// We could even in-line the func as anonymous/lambda
	fmt.Println(doMath(add, 1, 2, 3, 4, 5))
	fmt.Println(doMath(sub, 1, 2, 3, 4, 5))
	fmt.Println(doMath(func(x, y int) int { return x + y }, 6, 7, 8, 9, 0))

	inc1 := incrementor()
	fmt.Println(inc1()) // Notice how the value increments because x is the same variable
	fmt.Println(inc1())
	fmt.Println(inc1())

	inc2 := incrementor() // This is a different func, so x is different
	fmt.Println(inc2())   // Notice how x "resets" to 1 because inc2() is a different scope to inc1()
	fmt.Println(inc2())
	fmt.Println(inc2())

	fmt.Println(factorial(5))
}

// func (r receiver) identifier(parameters) (returns) { ... }

// We can do recursion with functions
func factorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	}

	return x * factorial(x-1)
}

func incrementor() func() int {
	var x int

	// The variable x is also scoped to this embedded func/code block
	return func() int {
		x++
		return x
	}
}

func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

// f is a callback function that takes in two parameters (int, int) and returns an int
func doMath(f func(x, y int) int, xi ...int) (res int) {
	res = 0

	for _, i := range xi {
		res = f(res, i)
	}

	return
}

// funcReturner returns a func(int, int) int
func funcReturner() func(x, y int) int {
	return func(x, y int) int { return x << y }
}

func doHumanStuff(h human) {
	// Assert that we're operating on a concrete type
	switch h := h.(type) {
	case person:
		fmt.Println("Doing human stuff for a normal person...")
		h.speak()
	case agent:
		fmt.Println("Doing human stuff for an agent...")
		h.speak()
	}
}

// The speak() function is attached to the receiver type of person
// Any value of type person can now make calls to speak()
func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am a mere person")
}

func (a agent) speak() {
	fmt.Println("May name is", a.first, a.last, "and I am an agent")
}

func (a agent) kill() {
	if a.ltk {
		fmt.Println("Bang! you're dead")
	} else {
		fmt.Println("I'm afraid I can't do that")
	}
}

// x is a type of []int
// A variadic parameter (...T) is just a slice ([]T)
// A variadic parameter can take in zero or more arguments
// If a func has a variadic parameter, then it must be the last parameter
func sumInts(x ...int) (sum int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum = 0

	for _, n := range x {
		sum += n
	}

	return // We can use simple return here since sum is our return value's identifier
	//return sum
}

func foo() {
	fmt.Println("Hello from foo")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println(s)
}

func woo(s string) int {
	return len(s) << 2
}

func revCat(fn, ln string) string {
	return fmt.Sprint(ln, ", ", fn)
}
