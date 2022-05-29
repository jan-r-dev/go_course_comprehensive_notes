package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()

	// loops over a channel until the channel is closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("exiting")
}
