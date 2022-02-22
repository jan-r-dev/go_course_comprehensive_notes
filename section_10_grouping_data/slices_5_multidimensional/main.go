package main

import "fmt"

func main() {
	jr := []string{"John", "Rig", "Vanilla", "Caramel"}
	mf := []string{"Martin", "French", "Raspberry", "Peach"}

	// Multi-dimensional slice
	p := [][]string{jr, mf}

	fmt.Println(p)
}
