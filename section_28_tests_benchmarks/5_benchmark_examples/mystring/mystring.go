package mystring

import "strings"

func Concat(xs []string) string {
	s := ""
	for i, v := range xs {
		if i == len(xs)-1 {
			s += v
			break
		}
		s += v
		s += " "
	}

	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}

/*
go test -bench .

goos: darwin
goarch: arm64
pkg: github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/5_benchmark_examples/mystring
BenchmarkConcat-10         92988             12872 ns/op
BenchmarkJoin-10          388002              3072 ns/op
PASS
ok      github.com/jan-r-dev/go_course_full_mcleod_notes/section_28_tests_benchmarks/5_benchmark_examples/mystring      2.723s

*/

/*
strings.Join()

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var b Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

*/
