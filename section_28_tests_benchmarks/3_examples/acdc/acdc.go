//Package acdc is here to sum your digits and rock your world.
package acdc

// MySum takes in a variadic number of integers and returns their sum.
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
