package main

import "fmt"

func main() {
	is := []int{5, 5, 5, 5}

	fmt.Println(
		foo(is...),
		bar(is),
	)
}

func foo(is ...int) int {
	total := 0
	for _, v := range is {
		total += v
	}

	return total
}

func bar(is []int) int {
	total := 0
	for _, v := range is {
		total += v
	}

	return total
}
