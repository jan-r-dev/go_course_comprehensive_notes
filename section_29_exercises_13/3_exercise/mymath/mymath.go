// Package mymath provides a function for calculating a centered average.
package mymath

import (
	"sort"
)

// CenteredAvg computes the average of a list of numbers after removing the largest and smallest values.
// If the slice contains a single element, it returns that element. If the slice contains two elements, it returns their simple average.
func CenteredAvg(xi []int) float64 {
	if len(xi) == 0 || len(xi) == 1 {
		return float64(xi[0])
	} else if len(xi) == 2 {
		return float64(xi[0]+xi[1]) / 2
	}

	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
