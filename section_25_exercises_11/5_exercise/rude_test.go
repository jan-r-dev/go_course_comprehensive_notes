package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	if got == "Go away" {
		t.Errorf("\nGreeting = %v;\nwant a welcome", got)
	}
}

/*

Common approach to testing in Go:
1/ We create a struct to represent the inputs and outputs for the function.
2/ Then we create a list of these structs (pairs).
3/ Then we loop through each one and run the function.

type testpair struct {
  values []float64
  average float64
}

var tests = []testpair{
  { []float64{1,2}, 1.5 },
  { []float64{1,1,1,1,1,1}, 1 },
  { []float64{-1,1}, 0 },
}

func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}
*/
