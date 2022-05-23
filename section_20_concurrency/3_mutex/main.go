package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// Lock started
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			// Lock ended - No other goroutine can read/write to the counter variable while the lock persists
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}
