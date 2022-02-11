package main

import "fmt"

// Bit shifting operators >> <<

func main() {

	kb := 1024
	mb := kb * 1024
	gb := mb * 1024

	fmt.Printf("%d\n%b\n", kb, kb)
	fmt.Printf("%d\n%b\n", mb, mb)
	fmt.Printf("%d\n%b\n", gb, gb)

}
