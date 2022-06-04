// Package myecho implements a way of printing multiple strings easily.
package myecho

import "fmt"

// Echo takes an unlimited number of strings and prints them all out on a separate line.
// This is done to demonstrate how the documentations can be generated.
func Echo(ss ...string) {
	for v, _ := range ss {
		fmt.Println(v)
	}
}

// Try:
// go doc
// go doc Echo
