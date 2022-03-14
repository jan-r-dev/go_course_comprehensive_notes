package main

import "fmt"

func main() {
	s1 := struct {
		colours     []string
		association map[string]string
	}{
		colours: []string{
			"blue",
			"green",
			"silver",
		},
		association: map[string]string{
			"car":   "Volkswagen",
			"money": "options",
			"books": "backlog",
		},
	}

	fmt.Println(s1)
	fmt.Println()

	for _, v := range s1.colours {
		fmt.Println(v)
	}

	for k, v := range s1.association {
		fmt.Printf("\t%v\t%v\n", k, v)
	}
}
