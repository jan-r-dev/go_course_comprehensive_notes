package main

import "fmt"

// Strings

func main() {
	s := "Hello there Ř"

	fmt.Printf("%T\n", s)
	fmt.Println(len(s))

	// Conversion to a byte slice
	bs := []uint8(s)
	fmt.Println(bs)
	fmt.Printf("%T\n\n\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
}
