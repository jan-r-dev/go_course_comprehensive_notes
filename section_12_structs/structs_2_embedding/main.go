package main

import "fmt"

// Composite data structure
type person struct {
	first string
	last  string
	age   int
}

// Includes an embedded struct
type secretAgent struct {
	person
	ltk bool
}

func main() {

	// Creating an embed
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	fmt.Println(sa1)

	// If no name collision exists, it is not necessary to specify sa1.person.first
	// The inner type person gets promoted to the outer type secretAgent
	fmt.Println(sa1.ltk, sa1.first, sa1.last, sa1.age)

}
