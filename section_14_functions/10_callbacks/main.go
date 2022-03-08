package main

import "fmt"

func main() {
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	totalAll := sum(is...)
	fmt.Println("Sum of all:\t", totalAll)

	totalEven := even(sum, is...)
	fmt.Println("Sum of even:\t", totalEven)

	totalOdd := odd(sum, is...)
	fmt.Println("Sum of even:\t", totalOdd)
}

func sum(is ...int) int {
	total := 0

	for _, v := range is {
		total += v
	}

	return total
}

func even(sumf func(is ...int) int, is2 ...int) int {
	ise := []int{}

	// This appends all the even numbers from is2 to ise (int slice even)
	for _, v := range is2 {
		if v%2 == 0 {
			ise = append(ise, v)
		}
	}

	totalEven := sumf(ise...)

	return totalEven
}

func odd(sumf func(is ...int) int, is2 ...int) int {
	iso := []int{}

	// This appends all the odd numbers from is2 to iso (int slice odd)
	for _, v := range is2 {
		if v%2 != 0 {
			iso = append(iso, v)
		}
	}

	totalOdd := sumf(iso...)

	return totalOdd
}
