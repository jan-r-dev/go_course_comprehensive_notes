package main

import "fmt"

type customErr struct {
	message string
}

func (c *customErr) Error() string {
	return fmt.Sprintf("Error occurred in a custom place: %v", c.message)
}

func main() {
	cerr := customErr{
		message: "some of the relevant things have broken",
	}

	foo(&cerr)

}

func foo(err error) {
	fmt.Println(err, err.(*customErr).message) // Example of assertion at the end
}
