package main

import "testing"

func TestMySum(t *testing.T) {
	const msg string = "expected %v, got %v"
	if v := mySum(); v != 0 {
		t.Errorf(msg, 0, v)
	}

	if v := mySum(2, 3); v != 5 {
		t.Errorf(msg, 5, v)
	}

	if v := mySum(1, 4); v != 5 {
		t.Errorf(msg, 5, v)
	}

	if v := mySum(6, 7); v != 13 {
		t.Errorf(msg, 13, v)
	}
}
