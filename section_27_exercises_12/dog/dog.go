// Package dog exports the Years function for converting human years into dog years.
package dog

// Years accepts human years as an integer and return dog years.
// Dog years are calculated as human years * 7
func Years(hy int) int {
	return hy * 7
}
