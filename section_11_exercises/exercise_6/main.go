package main

import "fmt"

func main() {
	x := make([]string, 0, 54)
	fmt.Printf("Length: %v\tCapacity: %v\n", len(x), cap(x))

	s := []string{"Alabama", "Arkansas", "American Samoa", "Arizona", "California", "Colorado", "Connecticut", "District of Columbia", "Delaware", "Florida", "Georgia", "Guam", "Hawaii", "Iowa", "Idaho", "Illinois", "Indiana", "Kansas", "Kentucky", "Louisiana", "Massachusetts", "Maryland", "Maine", "Michigan", "Minnesota", "Missouri", "Mississippi", "Montana", "North Carolina", "North Dakota", "Nebraska", "New Hampshire", "New Jersey", "New Mexico", "Nevada", "New York", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Puerto Rico", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Virginia", "Virgin Islands", "Vermont", "Washington", "Wisconsin", "West Virginia", "Wyoming"}

	x = append(x, s...)

	for i := 0; i < len(x); i++ {
		fmt.Printf("Index: %v\tState: %v\n", i, x[i])
	}

	fmt.Printf("Length: %v\tCapacity: %v\n", len(x), cap(x))
}
