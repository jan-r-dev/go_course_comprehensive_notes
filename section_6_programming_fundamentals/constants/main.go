package main

import "fmt"

/*
const a = 42
const b = 42.78
const c = "Samwise"

OR
*/

const (
	// Typed constant - no flexibility for the compiler
	a int = 42
	// Untyped constant
	b = 42.78
	c = "Samwise"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n%T\n%T\n", a, b, c)

}
