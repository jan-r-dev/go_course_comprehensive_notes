package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	// tests is a slice of multiple input + answer conditions for the test to go over
	tests := []test{
		{[]int{12, 32}, 44},
		{[]int{2, 2, 2, 0}, 6},
		{[]int{-2, 2, 5}, 5},
	}

	for _, v := range tests {
		r := mySum(v.data...)
		if r != v.answer {
			t.Error("Expected", v.answer, "Got", r)
		}
	}
}
