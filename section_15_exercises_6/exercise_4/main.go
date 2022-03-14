package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Name: %v %v\nAge: %v\n", p.first, p.last, p.age)
}

func main() {
	commisar := person{
		first: "Ciaphas",
		last:  "Cain",
		age:   65,
	}

	commisar.speak()
}
