package main

import "fmt"

func main() {
	f := foo()

	f()
	f()
	f()
	f()

	// Outside of closure: fmt.Println(x)
}

func foo() func() {
	x := 0

	return func() {
		x++
		fmt.Println(x)
	}
}
