package main

import "fmt"

func main() {

	deferMe()

}

func deferMe() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
