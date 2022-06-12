package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

/////

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

/////

func TestYears(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := Years(i)
		if res != i*7 {
			t.Error("got", res, "expected", i*7)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := YearsTwo(i)
		if res != i*7 {
			t.Error("got", res, "expected", i*7)
		}
	}
}

/*

go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/1_exercise/dog
BenchmarkYears-10               1000000000               0.3763 ns/op
BenchmarkYearsTwo-10            228669792                5.261 ns/op
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/1_exercise/dog     2.310s
go test
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/1_exercise/dog     0.075s

*/
