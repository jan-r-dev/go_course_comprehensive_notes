package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	p := person{
		age:  75,
		name: "Jeff",
	}
	fmt.Println(p.age)

	changeMe(&p)
	fmt.Println(p.age)

}

func changeMe(p *person) {
	p.age = 42
	// (*p).age = 232 This also works
}
