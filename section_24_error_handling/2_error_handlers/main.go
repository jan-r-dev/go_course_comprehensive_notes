package main

import (
	"fmt"
	"log"
	"os"
)

/*
Error-handling options

fmt.Println()
	Goes to stdout
log.Println()
	Includes timestamp; Goes to stdout by default; can go to files, wherever
log.Fatalln()
	Equivalent to Println followed by os.Exit()
log.Panicln()
	deferred functions run
	can use “recover”
	Once all goroutines have terminated or the panic has been recovered, the value of the argument supplied to panic is returned + error
panic()
*/

func main() {
	// Set destination for logger
	f, err := os.Create("main.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("not-exist.txt")
	if err != nil {
		fmt.Println("Error logged")
		log.Println(err)
		// log.Fatalln(err)
		// log.Panicln(err)
		// panic(err)
	}
	defer f2.Close()
}
