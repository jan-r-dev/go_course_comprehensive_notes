package main

import "fmt"

func main() {
	switch {
	case 4 == 4:
		fmt.Println("True with fallthrough")
		// fallthrough causes the case below to execute as well, even if it is false
		fallthrough
	case 2 == 3:
		fmt.Println("False")
		fallthrough
	case false:
		fmt.Println("False 2")
	case 3 == 4:
		fmt.Println("Yes3")
	case !(5 == 5):
		fmt.Println("Yes4")
	// Default only executes if none of the other cases match
	default:
		fmt.Println("default")
	}

	n := "Jan"
	switch n {
	// Multiple options in a case
	case "Pavel", "Jan":
		fmt.Println("Jan or Pavel")
	}
}
