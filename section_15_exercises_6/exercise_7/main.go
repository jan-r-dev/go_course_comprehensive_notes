package main

import "fmt"

func main() {
	fn := func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	}

	fn()
	fmt.Printf("Type of fn is: %T\n", fn)
}
