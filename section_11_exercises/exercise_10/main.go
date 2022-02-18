package main

import "fmt"

func main() {
	m := map[string][]string{
		"franz_karl":        {"Hammers", "Elector counts", "Marienburg"},
		"toddbringer_boris": {"Middenheim", "Being a jerk"},
		"von_carstein_vlad": {"Blood", "More blood", "Not lasting longer than 2 turns"},
	}

	m["settra"] = []string{"not serving", "being invincible", "banishing Khatep"}

	delete(m, "toddbringer_boris")

	for k, v := range m {
		fmt.Printf("This is: %v\n", k)
		for i, v := range v {
			fmt.Printf("\t%v\t%v\n", i, v)
		}
	}

}
