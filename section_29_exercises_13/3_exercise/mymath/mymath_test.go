package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{10, 17, 96, 57, 65, 50, 39, 37, 41, 53})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4}))
	// Output:
	// 2.5
}

func TestCenteredAvg(t *testing.T) {
	type test struct {
		values []int
		answer float64
	}

	tests := []test{
		{values: []int{1, 3, 4, 1}, answer: 2.0},
		{values: []int{1, 2, 4}, answer: 2.0},
		{values: []int{1, 2}, answer: 1.5},
		{values: []int{0}, answer: 0},
		{values: []int{-2, 4, 5, 1}, answer: 2.5},
	}

	for _, v := range tests {
		res := CenteredAvg(v.values)
		if res != v.answer {
			t.Error("got", res, "expected", v.answer)
		}
	}
}

/*
mymath % go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/3_exercise/mymath
BenchmarkCenteredAvg-10          7739544               154.3 ns/op
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/3_exercise/mymath  1.530s
mymath % go test
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/3_exercise/mymath  0.080s
mymath % go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/3_exercise/mymath  0.077s
*/
