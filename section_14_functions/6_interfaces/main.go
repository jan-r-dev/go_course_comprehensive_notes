package main

import "fmt"

// Interfaces allow us to define behaviour
// Interfaces allow for polymorphism; for a value to be of more than one type

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last, "and an agent.")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and a person.")
}

func foo(h human) {
	h.speak()
	fmt.Println("I was passed into the human interface.")
}

// Examples of assertion
func bar(h human) {
	// Switch statement over the types
	switch h.(type) {
	case person:
		fmt.Println(h.(person).first)
	case secretAgent:
		fmt.Println(h.(secretAgent).first)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Ciaphas",
		last:  "Cain",
	}

	// type secretAgent as argument; also is of type human
	foo(sa1)
	// type person as argument; also is of type human
	foo(p1)

	// Assertion
	bar(sa1)
	bar(p1)
}
