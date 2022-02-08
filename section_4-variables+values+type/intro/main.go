package main

import "fmt"

// Declaring a variable outside of a function body using the var keyword
// Assigns the value of 20
var y = 20

// Declaring a variable with the identifier 'z' and of a type int
// Assigns the zero value of type int to z
var z int

// Declaring a variable with the identifier ss of a type string with the value Hello
// Once created, it impossible to change the contents of a string
var ss string = "Hello"

// Raw string literal - keeps everything except backticks
var bt string = `hi

there`

func main() {

	n, _ := fmt.Println("Hello there")
	fmt.Println(n)

	// Declare a variable and assign a value of a certain type
	// A statement
	x := 10

	fmt.Println(z)
	// An expression
	z = x + y

	fmt.Println(z, ss, bt)
}

/*
/ Every value is of type interface{} -- an empty interface
/ ...interface{} = give me as many values of type interface{} as you want
*/
