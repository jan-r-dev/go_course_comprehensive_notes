package main

import "testing"

func TestMySum(t *testing.T) {
	r := mySum(2, 3)
	if r != 5 {
		t.Error("Expected", 5, "Got", r)
	}
}
