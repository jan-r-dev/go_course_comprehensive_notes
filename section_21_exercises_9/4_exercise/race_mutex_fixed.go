package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	gr := 50
	counter := 0

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
