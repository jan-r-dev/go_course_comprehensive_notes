package main

import "fmt"

// Bit shifting operators >> <<

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\n%b\n", kb, kb)
	fmt.Printf("%d\n%b\n", mb, mb)
	fmt.Printf("%d\n%b\n", gb, gb)

}

/*
kb := 1024
mb := kb * 1024
gb := mb * 1024

*/
