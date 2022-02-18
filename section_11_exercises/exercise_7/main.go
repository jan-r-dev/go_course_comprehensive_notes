package main

import "fmt"

func main() {
	ss := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hellooo, James"},
	}

	for _, v := range ss {
		fmt.Printf("%T\n", v)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}
}
