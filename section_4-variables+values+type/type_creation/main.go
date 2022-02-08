package main

import "fmt"

var a int

// Declaring a custom type based on the type int
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Converts the value of type hotdog into value of type int
	a = int(b)
}
