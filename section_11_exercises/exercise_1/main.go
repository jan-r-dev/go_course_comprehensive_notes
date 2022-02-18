package main

import "fmt"

func main() {
	x := [5]int{}

	x[0] = 20
	x[1] = 5
	x[2] = 123
	x[3] = 2
	x[4] = 77

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("\n%T\n", x)

}
