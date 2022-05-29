package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go gensend(c)
	receive(c)

}

func gensend(c chan<- int) {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(c)
}

func receive(c <-chan int) {
	i := 0
	for v := range c {
		fmt.Println(v)
		i++
	}

	fmt.Println("----------")
	fmt.Println(i)
}
