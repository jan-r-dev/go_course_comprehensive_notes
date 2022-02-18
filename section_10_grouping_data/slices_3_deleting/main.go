package main

import "fmt"

func main() {
	x := []int{10, 20, 32, 12}
	y := []int{112, 2000, 56}

	x = append(x, y...)
	x = append(x, 11111, 23)
	fmt.Println(x)

	// Trim a slice to go from 0 to 2 and from 4 to end
	// Remember that append takes a slice of T and an unlimited amount of values of T
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
