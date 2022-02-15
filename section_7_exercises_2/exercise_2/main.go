package main

import "fmt"

func main() {
	a := 10 == 10
	b := 10 <= 9
	c := 10 >= 11
	d := 10 != 10
	e := 10 > 8
	f := 10 < 8

	fmt.Println(a, b, c, d, e, f)
}
