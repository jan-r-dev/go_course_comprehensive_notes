package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go foo("Hi there", &wg)
	go foo("Greetings", &wg)

	wg.Wait()
}

func foo(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}
