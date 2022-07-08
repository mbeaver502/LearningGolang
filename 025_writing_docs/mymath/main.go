// go doc mymath
// go doc mymath.Sum
// go doc mymath Sum

// Package mymath provides math solutions.
package mymath

// Sum adds an unlimited number of values of type int.
// We can have comments that need to go to a new line,
// and to do that we simply break where it makes sense.
func Sum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}
