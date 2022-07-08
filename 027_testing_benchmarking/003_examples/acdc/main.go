// Package acdc asks if you are ready to rock.
// If you are, then it salutes you. :)
package acdc

// Sum adds an unlimited number of values of type int.
func Sum(n ...int) int {
	sum := 0

	for _, i := range n {
		sum += i
	}

	return sum
}
