package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a) // &a gives the memory address of a

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // &a is a pointer to an int (*int)

	// Everything in Go is pass by value, but we can use pointers to pass addresses
	// Maybe we don't want to pass around some large data (maybe it's a huge struct)
	// We can use pointers to simplify that so we're not copying around all that data
	b := &a
	fmt.Println(b)  // b is now a pointer to a
	fmt.Println(*b) // dereference the pointer stored in b to get the stored value at that address

	*b = 50 // we can modify the dereferenced value, so we're actually updating the value stored by a
	fmt.Println(a)

	fmt.Println("------------")

	x := 0
	fmt.Println("address of x:", &x)
	foo(x) // x will be unmodified
	fmt.Println(x)

	fmt.Println("address of x:", &x)
	bar(&x) // x could be modified
	fmt.Println(x)

	// Method sets:
	// Method with non-pointer receiver works with both pointer and non-pointer values
	//  (t T) => T and *T
	// Method with pointer receiver works only with pointer values
	//  (t *T) => *T
}

// The value passed into here will not be modified
func foo(y int) {
	fmt.Println("address of y:", &y)
	fmt.Println(y)

	y = 43
	fmt.Println(y)
}

// We can use the pointer to modify the underlying value of y
func bar(y *int) {
	fmt.Println("address of y:", y)
	fmt.Println(*y)

	*y = 43
	fmt.Println(*y)
}
