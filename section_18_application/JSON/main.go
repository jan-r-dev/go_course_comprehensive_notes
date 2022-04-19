package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Jeff",
		Last:  "Bridges",
		Age:   12,
	}
	people := []person{p1, p2}

	// Marshal into a byte slice
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(string(bs))

	// Unmarshal into the slice of structs (using a pointer)
	p3 := []person{}
	err = json.Unmarshal(bs, &p3)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(p3)

}
