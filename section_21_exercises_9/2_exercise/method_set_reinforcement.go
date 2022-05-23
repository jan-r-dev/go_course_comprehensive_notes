package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

// Since the method is attached to a pointer, it will only accept a pointer
func (p *person) speak() {
	fmt.Println(p.name)
}

// Method accepting a type T would accept both a variable of type T and a pointer
/*
func (p person) speak() {
	fmt.Println(p.name)
}
*/

func main() {
	p1 := person{
		name: "MethodSet",
	}

	// Below will work
	saySomething(&p1)
	// Below will not work
	//saySomething(p1)
}

func saySomething(h human) {
	h.speak()
}
