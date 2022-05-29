package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		// the sender blocks until a receiver has received the value
		c <- 42
	}()

	// the receiver blocks until it has received the value from a sender
	fmt.Println(<-c)
}
