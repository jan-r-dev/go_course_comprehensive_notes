package main

import "fmt"

// Tests are in files that end with _test.go
// Tests are in the same package as the file being tested
// Tests are in a func with a signature "func TestXxx(*testing.T)"
// Command: go test

func main() {
	fmt.Println(mySum(1, 2, 3, 4))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// To run test: go test -v (-v flag for verbose)
