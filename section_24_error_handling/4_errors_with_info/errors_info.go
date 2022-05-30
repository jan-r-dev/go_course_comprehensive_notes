package main

import (
	"fmt"
	"log"
)

type mathError struct {
	style     string
	operation string
	err       error
}

func (n mathError) Error() string {
	return fmt.Sprintf("Math error occurred: %v, %v, %v", n.style, n.operation, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Operation attempted on negative number error: %v", f)
		return 0, mathError{"Test", "Square root", nme}
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
