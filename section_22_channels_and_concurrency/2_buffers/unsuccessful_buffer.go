package main

import "fmt"

func main() {
	c := make(chan int, 1)

	// will not run. channel is buffered only for 1
	c <- 42
	c <- 43

	fmt.Println(<-c)
}
