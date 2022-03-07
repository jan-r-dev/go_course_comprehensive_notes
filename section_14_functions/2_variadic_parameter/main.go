package main

import (
	"fmt"
)

func main() {
	f1("Hello", "there", "General Kenobi")
}

// A variadic amount of parameters of type string
// Note that this parameter must be the last in the list
// (s .. string, i int) is not possible; (i int, s .. string) is fine
func f1(s ...string) {
	fmt.Printf("%T\t%v\n", s, s)

	s2 := ""

	for _, v := range s {
		s2 = s2 + " " + v
	}

	fmt.Println(s2)
}
