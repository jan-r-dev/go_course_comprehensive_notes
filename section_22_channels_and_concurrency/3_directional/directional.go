package main

import "fmt"

func main() {
	c := make(chan int)
	// Channel that can only send int
	cs := make(chan<- int)
	// Channel that can only receive int
	cr := make(<-chan int)

	// Read from left to right. <-chan receive from channel; chan<- am on channel, sending

	fmt.Printf("%T\n%T\n", cs, cr)

	// specific to general doesn't assign
	/*
		c = cr
		c = cs
	*/

	// general to specific does assign
	cr = c
	cs = c
}
