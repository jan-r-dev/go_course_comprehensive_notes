package main

import "fmt"

func main() {
	var favSport string = "Grgling"

	switch favSport {
	case "Skating":
		fmt.Println("Your favourite sport is skating.")
	case "Skiing":
		fmt.Println("Your favourite sport is skiing.")
	default:
		fmt.Println("What on earth are you doing")
	}

}
