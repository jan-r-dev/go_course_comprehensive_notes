package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome and greetings, James." {
		t.Error("got", s, "expected", "Welcome and greetings, James.")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome and greetings, James.
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}

// Benchmarking is done with i < b.N -> This means that the loop will run as many times as needed to get a statistically accurate number.
// Run benchmark: go test -bench . OR go test -bench GREET

// For coverage, run: go test -cover OR go test -coverprofile cover.out
// go test -coverprofile cover.out -> go tool cover -html=cover.out
