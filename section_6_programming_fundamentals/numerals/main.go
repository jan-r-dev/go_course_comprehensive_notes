package main

import "fmt"

func main() {
	// Converting H into a byte and then into various numeral systems
	s := "H"

	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	// In decimal base10
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	// In binary base2
	fmt.Printf("%b\n", n)

	// In hex base16
	fmt.Printf("%x\n", n)
	// In hex base16 with 0X notation
	fmt.Printf("%#X\n", n)
}

//
