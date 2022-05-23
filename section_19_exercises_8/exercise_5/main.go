package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (users ByAge) Len() int           { return len(users) }
func (users ByAge) Swap(i, j int)      { users[i], users[j] = users[j], users[i] }
func (users ByAge) Less(i, j int) bool { return users[i].Age < users[j].Age }

type ByLast []user

func (users ByLast) Len() int           { return len(users) }
func (users ByLast) Swap(i, j int)      { users[i], users[j] = users[j], users[i] }
func (users ByLast) Less(i, j int) bool { return users[i].Last < users[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	for _, user := range users {
		sort.Strings(user.Sayings)
	}
	printUsers(users)

	sort.Sort(ByAge(users))
	printUsers(users)

	sort.Sort(ByLast(users))
	printUsers(users)
}

func printUsers(users []user) {
	for _, user := range users {
		fmt.Println(user.First, user.Last)
		fmt.Println("Age:", user.Age)
		for _, saying := range user.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
	}
	fmt.Println("======================================")
}
