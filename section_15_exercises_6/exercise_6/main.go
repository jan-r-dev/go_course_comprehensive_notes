package main

import "fmt"

func main() {
	func(i int) {
		fmt.Println("Anonymously int:", i)
	}(5)
}
