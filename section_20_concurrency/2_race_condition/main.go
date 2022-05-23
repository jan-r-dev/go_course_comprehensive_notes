package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Race condition intro
func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// Yields to CPU, allowing something else to run
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	// fmt.Println("GoRoutines\t", runtime.NumGoroutine()) -- executing it here would let us see
	// the other GoRoutines before they were finished
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	// go run -race main.go can detect race conditions
}
