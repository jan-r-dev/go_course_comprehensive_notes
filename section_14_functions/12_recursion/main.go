package main

import "fmt"

func main() {
	n := factorialRec(4)
	n2 := factorialLoop(4)

	fmt.Println(n)
	fmt.Println(n2)
}

func factorialRec(n int) int {
	if n == 0 {
		return 1
	}
	return n * (factorialRec(n - 1))

	// n * (factorial(4 - 1)) * (factorial(3 - 1)) * (factorial(2 - 1)) * 1
}

func factorialLoop(n int) int {
	total := 1

	for ; n > 0; n-- {
		total *= n
	}

	return total
}
