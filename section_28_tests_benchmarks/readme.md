Remember to BET

- Benchmark
- Example
- Test

```
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome and greetings, James.
}

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome and greetings, James." {
		t.Error("got", s, "expected", "Welcome and greetings, James.")
	}
}
```

Useful commands:

```
godoc -http=:8080
go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```