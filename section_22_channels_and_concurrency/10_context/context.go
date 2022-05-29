package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Printf("%T\n", ctx)
	fmt.Printf("%T\n", cancel)
	fmt.Println()

	fmt.Println("Error:", ctx.Err())
	fmt.Println("Goroutine count:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			// ctx.Done() is provided for use in select statements
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error:", ctx.Err())
	fmt.Println("Goroutine count:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("Error:", ctx.Err())
	fmt.Println("Goroutine count:", runtime.NumGoroutine())
}
