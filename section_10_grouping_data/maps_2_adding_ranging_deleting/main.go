package main

import "fmt"

func main() {
	m := map[string]string{
		"Karl":  "Franz",
		"Felix": "Jaeger",
	}
	fmt.Println(m)

	// Adding a new key-value pair to a map
	m["Boris"] = "Toddbringer"
	fmt.Println(m)

	// Deleting a key-value pair from a map
	delete(m, "Felix")

	// Ranging over the key-value pairs in a map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// If the key-value pair exists, delete it
	if v, ok := m["Karl"]; ok {
		fmt.Println(v)
		delete(m, "Karl")
	}

	fmt.Println(m)
}
