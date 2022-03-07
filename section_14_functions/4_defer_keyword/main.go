package main

import "fmt"

func main() {
	foo()
}

func foo() {
	defer bar()

	fmt.Printf("Starting... ")
}

func bar() {
	fmt.Println("and ending")
}
