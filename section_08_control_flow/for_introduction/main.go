package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i < 3; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i)
			fmt.Print(j)
		}
		fmt.Println()
	}

	fmt.Println("-------------------")
	i := 1
	for ; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------------")
	j := 1
	for {
		// Printing only even numbers
		j++
		if j > 100 {
			break
		}

		if j%2 != 0 {
			continue
		}

		fmt.Println(j)
	}
}
