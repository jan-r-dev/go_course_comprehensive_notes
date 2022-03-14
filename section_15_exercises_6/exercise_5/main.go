package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	c := circle{
		radius: 10,
	}

	s := square{
		side: 10,
	}

	info(c)
	info(s)
}

func info(s shape) {
	fmt.Println(
		"Area:",
		// Rounding to 3 decimals
		(math.Round(s.area()*1000) / 1000),
	)
}
