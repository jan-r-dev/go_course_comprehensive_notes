package main

import "fmt"

const (
	// Automatically incremented by 1
	a = iota
	b
	c
)

const (
	// Resets to 0 with every const block
	d = iota
	e = iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
