package main

import "fmt"

func main() {
	f := foo()
	f()

	// Returns a function that will return an int
	f2 := bar()
	// Returns an int
	i := f2()
	fmt.Println("Returned function returning int:", i)

	// Alternate execution 1
	f3 := bar()
	fmt.Println("Alternatively 1: ", f3())

	// Alternate execution 2
	fmt.Println("Alternatively 2: ", bar()())
}

func foo() func() {

	return func() {
		fmt.Println("Returned function")
	}
}

func bar() func() int {

	return func() int {
		return 9
	}
}
