package main

import "fmt"

func main() {
	// Functions called with arguments
	s2, b := foo("hello")
	fmt.Println(s2, "2", b)
}

// Functions defined with parameters
func foo(s string) (string, bool) {
	return fmt.Sprint(s), true
}

// Pass by value
// Pass by value means that a copy of the actual parameter's value is made in memory, i.e. the caller and callee have two independent variables with the same value. If the callee modifies the parameter value, the effect is not visible to the caller.
