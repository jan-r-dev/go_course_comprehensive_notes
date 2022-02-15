package main

import "fmt"

var i int

func main() {
	i = 52

	fmt.Printf("%d\n%b\n%#x\n", i, i, i)

	bi := i << 1

	fmt.Printf("%d\n%b\n%#x\n", bi, bi, bi)
}
