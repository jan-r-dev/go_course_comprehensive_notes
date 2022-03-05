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

	pm := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	// Using pm[k] instead of the value for fun
	for k, _ := range pm {
		fmt.Println(pm[k].first_name, pm[k].last_name)

		fmt.Println("Favourite flavours:")
		for _, v := range pm[k].favourite_flavours {
			fmt.Printf("\t%v\n", v)
		}
	}

}
