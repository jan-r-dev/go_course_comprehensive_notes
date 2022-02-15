package main

import "fmt"

// Use format printing to loop over numbers 0 to 100 and print them as letters in ascii

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Printf("%c\n", i)
	}

}
