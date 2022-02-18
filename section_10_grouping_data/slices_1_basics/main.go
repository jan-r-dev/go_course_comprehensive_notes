package main

import "fmt"

func main() {
	// composite literal syntax: x := type{values}
	// Creating a slice
	x := []int{10, 20, 32, 12}

	// Examining a slice
	fmt.Printf("Type: %T\tLength: %v\tCap: %v\tContent: %v\n", x, len(x), cap(x), x)

	// Ranging over a slice, using index and value
	for i, v := range x {
		fmt.Println(i, v)
	}
}
