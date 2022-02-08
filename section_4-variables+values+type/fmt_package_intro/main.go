package main

// https://pkg.go.dev/fmt
import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	// Value in default format
	fmt.Printf("%v\n", y)
	// Type of value
	fmt.Printf("%T\n", y)

	// Binary
	fmt.Printf("%b\n", y)
	// Hexadecimal
	fmt.Printf("%x\n", y)
	// Hexadecimal with leading 0x
	fmt.Printf("%#x\n", y)

	// Combination separated by tabs. Escape characters like tabs and newlines can be found in the language specification
	fmt.Printf("%#x\t%b\t%T\n", y, y, y)

	// Sprint assigns the result to a string
	s := fmt.Sprintf("%#x\t%b\t%T\n", y, y, y)

	fmt.Println(s)
}
