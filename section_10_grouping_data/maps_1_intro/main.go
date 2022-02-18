package main

import "fmt"

// maps
// efficient and fast look-up
// unordered list

func main() {
	// Key: string Value: int
	m := map[string]int{
		"James": 32,
		"John":  12,
	}

	fmt.Println(m)

	fmt.Println(m["James"])
	// If a key is not found in the map, it returns a zero value for that type
	fmt.Println(m["Random"])

	// Value: 0, ok: false
	v, ok := m["Random"]
	fmt.Println(v, ok)

	if v, ok := m["James"]; ok {
		fmt.Printf("Age of james: %v. It is ok.\n", v)
	}
}
