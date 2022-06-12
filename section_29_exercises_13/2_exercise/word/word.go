// Package word helps with counting the number and occurences of words in strings.
package word

import (
	"strings"
)

// UseCount takes a string and returns a map of all whitespace-delimited words with the number of times that they occur in the string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range xs {
		m[v]++
	}

	return m
}

// Count returns the number of words in a string separated by one or more consecutive whitespace characters
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
