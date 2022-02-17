package main

import "fmt"

func main() {
	// Declaring an array like this requires specifying the size
	// Fills in zero values
	var x [5]int

	x[3] = 42

	fmt.Println(x)
	// Length is part of the array's type
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

	// General conclusion: Use slices instead
}
