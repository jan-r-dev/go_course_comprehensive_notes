package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	gr := 50
	counter := 0

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
