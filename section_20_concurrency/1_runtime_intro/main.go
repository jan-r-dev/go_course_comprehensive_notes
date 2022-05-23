package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())

	// Specifying that we will wait for two things
	wg.Add(2)
	go foo()
	bar()

	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	// Waiting for wg.Done()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Foo:", i)
	}
	// One thing finishes
	wg.Done()
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Bar:", i)
	}
	// One thing finishes
	wg.Done()
}

/*

Parallelism
	Based on multiple threads running at the same time (requires multiple logical CPUs)

Concurrency
	Design pattern
	Has the potential to run in parallel
	Does not guarantee parallelism
*/
