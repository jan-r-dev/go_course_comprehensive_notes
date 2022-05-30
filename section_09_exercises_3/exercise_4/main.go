package main

import "fmt"

func main() {
	x := 1994

	for {
		fmt.Println(x)
		x++

		if x > 2022 {
			break
		}
	}

}
