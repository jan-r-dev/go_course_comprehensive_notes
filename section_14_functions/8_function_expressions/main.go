package main

import "fmt"

func main() {
	// Assigning a function to a variable is what is known as a function expression
	// Functions in Go are first-class citizens -- they are a type as any other type
	f := func() {
		fmt.Println("Expressive.")
	}

	f()

	f2 := func(x int) {
		fmt.Println(x)
	}

	f2(24)
}
