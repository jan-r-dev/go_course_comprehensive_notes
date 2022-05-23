package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "134"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	pass := "134"
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil {
		log.Fatal("Passwords do not match")
	}
	fmt.Println("Logged in")
}
