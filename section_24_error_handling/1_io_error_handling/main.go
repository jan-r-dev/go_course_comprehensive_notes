package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
type error is an easily implementable interface

type error interface {
	Error() string
}
*/

func main() {
	createAndWrite()
	openToRead()
}

func createAndWrite() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Joshua")

	io.Copy(f, r)

}

func openToRead() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
