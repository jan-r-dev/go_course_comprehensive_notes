package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	// Unfluring the []int into a number of arguments of type int
	// Note that unfurling just means the compiler is informed that the argument is a slice of the type requred
	// In this case no new slice is created by the function call
	f1(ii...)

	// Variadic parameter means zero arguments is also an option
	f1()
}

// A variadic amount of parameters of type string
func f1(i ...int) {

	sum := 0

	for _, v := range i {
		sum = sum + v
	}

	fmt.Println(sum)
}
