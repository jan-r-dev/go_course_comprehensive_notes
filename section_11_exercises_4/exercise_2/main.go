package main

import "fmt"

func main() {
	x := []int{20, 5, 123, 2, 77}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("\n%T\n", x)

}
