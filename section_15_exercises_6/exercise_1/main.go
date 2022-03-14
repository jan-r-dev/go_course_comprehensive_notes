package main

import "fmt"

func main() {
	i := foo()
	i2, s := bar()

	fmt.Println(i)
	fmt.Println(i2, s)
}

func foo() int {
	i := 1 + 4

	return i
}

func bar() (int, string) {
	i := 3
	s := "Hello"

	return i, s
}
