package main

import "fmt"

func main() {
	if true {
		fmt.Println(1)
	} else if !false {
		fmt.Println(4)
	} else if !(2 == 2) {
		fmt.Println(5)
	} else if 2 == 2 {
		fmt.Println(6)
	}

	// Initialisation statement
	// Limits the scope of the variable x to the if statement
	if x := 42; x == 2 {
		fmt.Println(x)
	}

	// Will not work; x is undefined outside of the scope of the if statement
	//fmt.Println(x)
}
