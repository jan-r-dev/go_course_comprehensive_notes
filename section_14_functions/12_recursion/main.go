package main

import "fmt"

var x int = 10
var y int = 1

func main() {
	recursor()

	fmt.Println(y)
}

func recursor() {
	if x != 1 {
		y += y * (x - 1)
		x--
		recursor()
	}
}

// Strangely proud of this one. Figured out solo without understanding recursion or factorials prior to trying.
