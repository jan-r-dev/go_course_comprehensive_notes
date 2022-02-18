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
	y := []int{112, 2000, 56}
	// Appending one slice to the other with unfurling the slice
	x = append(x, y...)
	// Appending values of the same type to the slice
	x = append(x, 11111, 23)

	fmt.Println(x)

	varyMe(x...)
}

// Variadic parameter ...int below. Takes unlimited number of values
func varyMe(y ...int) {
	for i, v := range y {
		fmt.Println(i, v)
	}
}
