package main

import (
	"fmt"

	"github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/3_examples/acdc"
)

func main() {
	r := acdc.MySum(1, 2, 23, 23)

	fmt.Println(r)
}

// To test everywhere, remember: go test ./...
// Means everything in the current folder and all the folders branching down

// Run godoc -http=:8080
// Access: http://localhost:8080/pkg/github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/3_examples/acdc/#pkg-examples

// Afterwards consult ./acdc/acdc_test.go
// Example code is ran and verified
