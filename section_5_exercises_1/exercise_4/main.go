package main

import "fmt"

type superint int

var x superint

func main() {
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
