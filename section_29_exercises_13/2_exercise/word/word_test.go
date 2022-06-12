package word

import (
	"fmt"
	"testing"

	"github.com/jan-r-dev/go_course_full_mcleod_notes/section_29_exercises_13/2_exercise/quote"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello there, general."))
	// Output:
	// 3
}

func TestCount(t *testing.T) {
	type test struct {
		data  string
		words int
	}

	tests := []test{
		{"", 0},
		{"Hello there, officer.", 3},
		{"With respect to the many instances...", 6},
		{quote.SunAlso, 1349},
	}

	for _, v := range tests {
		res := Count(v.data)
		if res != v.words {
			t.Error("got", res, "expected", v.words)
		}
	}

}

func TestUseCount(t *testing.T) {
	test := struct {
		data  string
		words map[string]int
	}{
		data: "Come slow now to everything. The heavy wing. The heavy wing. Wings and wing.",
		words: map[string]int{
			"Come": 1, "slow": 1, "now": 1, "to": 1, "everything.": 1, "The": 2, "heavy": 2, "wing.": 3, "Wings": 1, "and": 1},
	}

	m := UseCount(test.data)

	for k, v := range m {
		if v != test.words[k] {
			t.Error("Got", k, "with", v, "Expected", k, "with", test.words[k])
		}
	}
}

/*
word % go test
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/2_exercise/word    0.294s
word % go test
--- FAIL: TestUseCount (0.00s)
    word_test.go:64: Got The with 2 Expected The with 3
FAIL
exit status 1
FAIL    github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/2_exercise/word    0.167s
word % go test
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/2_exercise/word    0.177s
word % go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/2_exercise/word
BenchmarkCount-10                  25646             45636 ns/op
BenchmarkUseCount-10                9343            120361 ns/op
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/section_29_exercises_13/2_exercise/word    2.941s
*/
