package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First       string
	Last        string
	Sayings     []string
	ChanToError chan int
}

func main() {
	p1 := person{
		First:       "James",
		Last:        "Bond",
		Sayings:     []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
		ChanToError: make(chan int),
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("Error occured in JSON marshaller:\t%v", err)
	}

	return bs, err
}
