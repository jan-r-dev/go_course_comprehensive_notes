package main

import "fmt"

func main() {

	n, _ := fmt.Println("Hello there")
	fmt.Println(n)

	x := 10 // Statement
	y := 20
	z := x + y // Expression

	fmt.Println(z)
}

/*
/ Every value is of type interface{} -- an empty interface
/ ...interface{} = give me as many values of type interface{} as you want
*/
