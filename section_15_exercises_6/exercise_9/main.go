package main

import "fmt"

func main() {
	i := foo(bar)

	fmt.Println(i)
}

func foo(f func() int) int {
	return f()
}

func bar() int {
	return 42
}
