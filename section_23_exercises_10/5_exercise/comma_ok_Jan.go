package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if v, ok := <-c; ok {
			fmt.Println(v, ok)
		} else {
			fmt.Println("Comma ok signal, exiting;", ok)
			return
		}
	}
}
