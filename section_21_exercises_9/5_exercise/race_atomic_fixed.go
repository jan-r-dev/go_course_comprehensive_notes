package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	gr := 50
	var counter int64

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			runtime.Gosched()

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
