package main

import "fmt"

type person struct {
	first_name         string
	last_name          string
	favourite_flavours []string
}

func main() {

	p1 := person{
		first_name:         "Jeff",
		last_name:          "Johnson",
		favourite_flavours: []string{"vanilla", "caramel"},
	}

	p2 := person{
		first_name:         "Adam",
		last_name:          "Adamite",
		favourite_flavours: []string{"chocolate", "pudding", "strawberry"},
	}

	ps := []person{p1, p2}

	for _, v := range ps {
		fmt.Println(v.first_name, v.last_name)
		fmt.Println("Favourite flavours:")

		for _, v := range v.favourite_flavours {
			fmt.Printf("\t%v\n", v)
		}
	}

}
