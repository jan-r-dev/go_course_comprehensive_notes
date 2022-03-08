package main

import "fmt"

func main() {

	{
		y := 42
		fmt.Println(y, "in a code block.")
	}

	// Cannot do - y undefined in this scope
	// fmt.Println(y)

	//

	a := incrementer()
	b := incrementer()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(a())

}

func incrementer() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
