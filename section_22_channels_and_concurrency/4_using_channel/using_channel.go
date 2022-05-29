package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)
	//receive (and block, since no go keyword; channel there blocks this goroutine until it receives)
	bar(c)

	fmt.Println("exiting")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
