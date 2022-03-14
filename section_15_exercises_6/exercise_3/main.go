package main

import "fmt"

func main() {
	defer lately()

	fmt.Println("Beginning")
}

func lately() {
	defer func() {
		fmt.Println("Ending for real.")
	}()

	fmt.Println("Ending")
}
