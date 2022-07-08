package main

import "testing"

func TestMySum(t *testing.T) {
	const msg string = "expected %v, got %v"

	type test struct {
		data     []int
		expected int
	}

	tests := []test{
		{
			data:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			data:     []int{},
			expected: 0,
		},
		{
			data:     []int{1},
			expected: 1,
		},
		{
			data:     []int{2, 2},
			expected: 4,
		},
	}

	for _, test := range tests {
		if actual := mySum(test.data...); actual != test.expected {
			t.Errorf(msg, test.expected, actual)
		}
	}

}
