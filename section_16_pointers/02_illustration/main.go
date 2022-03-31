package main

import "fmt"

func main() {
	x := 0
	foo(x)
	fmt.Println(x)

	fmt.Println("===========================")

	x = 0
	bar(&x)
	fmt.Println(x)

}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}
