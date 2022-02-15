package main

import "fmt"

const (
	y1 = 2019
	y2 = y1 + iota
	y3 = y1 + iota
	y4 = y1 + iota
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}
