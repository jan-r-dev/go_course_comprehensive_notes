package main

import "fmt"

func main() {
	a := 1

	// Prints value
	fmt.Println(a)
	// Prints address in memory
	fmt.Println(&a)
	// Prints types of the variable int and the pointer *int
	fmt.Printf("%T\t%T\n", a, &a)

	b := &a
	// Using * as an operator. Dereferences an address -> Gives value stored at an address
	fmt.Println()
	fmt.Println(b)
	fmt.Println(*b)
	// & gives you the address. * gives you the value stored at the address
	fmt.Println(*&a)
	fmt.Println()

	// Modify the value found at the address. This means a also changes, since it occupies the same space in memory
	*b = 63
	fmt.Println(a)
}
