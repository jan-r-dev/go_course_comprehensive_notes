package main

import "fmt"

func main() {

	func(x int) {
		fmt.Println("I am anonymous func.")
		fmt.Println(x)
	}(42)
}
