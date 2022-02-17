package main

import "fmt"

func main() {
	x := []int{10, 20, 32, 12}

	// From position 1 to the end, including 1
	fmt.Println(x[1:])

	// From the begnning to position 2, excluding 2
	fmt.Println(x[:2])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Appending
	x = append(x, 1, 2, 3, 56)

	fmt.Println(x)
}
