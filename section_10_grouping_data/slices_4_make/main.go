package main

import "fmt"

func main() {
	x := make([]float64, 10, 12)
	fmt.Println(x, len(x), cap(x))

	// Attempting to allocate a value outside of the length results in:
	// panic: runtime error: index out of range [10] with length 10

	// Appending extends the length, and keeps the same capacity if sufficient
	x = append(x, 123)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 123)
	fmt.Println(x, len(x), cap(x))

	// Appending above the capacity causes the runtime to double the underlying array's capacity
	x = append(x, 123)
	fmt.Println(x, len(x), cap(x))
}
