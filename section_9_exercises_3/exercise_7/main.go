package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("Printing odd number: %v\n", i)
		} else if i == 10 {
			fmt.Println("End of loop")
		} else {
			fmt.Printf("Printing even number: %v\n", i)
		}
	}
}
