package main

import "fmt"

// Use format printing to loop over numbers 0 to 100 and print them as letters in ascii

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}

}
