package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		switch {
		case i == 1:
			fmt.Printf("1 == 1 is: %v\n", 1 == 1)
		case i == 2:
			fmt.Printf("1 == 2 is: %v\n", i == 2)
		case i == 1+2:
			fmt.Printf("3 == 1 + 2 is: %v\n", 3 == 1+2)
		}
	}

}
