package main

import "fmt"

func main() {
	fmt.Println("I require a new project to play with")
	fmt.Println("--------------------------------------")

	fmt.Println("Let's list the things I want to cover.")
	fmt.Println("test-driven development -> writing tests first", "")

	fmt.Printf("\n\n%v\n\n", "--------------------------------------")

	fmt.Println("Let's list the things I don't want to devote time for this one")
	fmt.Println("Extensive work on SQL")

}

/*

SSO
Crypto
Certificates
Authentication & Authorization
Docker & Podman
Something to make myself better

...

Something to make myself better ->
	Productivity tools for work.
	Which areas need to be more productive? ->
		Preparing all the commands for production takes a while
		That could be resolved by practice.
		It is also likely that many of the tasks will be relegated to Rundeck as well
		...
		Regular expressions are a weak-point of mine
		Does Go implement Perl-like regex?
		...
		It does ->
			https://pkg.go.dev/regexp
			...
			Possible

Something to build upon my current skillset ->
	Postgres integration with Go ->
		Current not working with dynamic structs, only hard-coded
		Unfinished
		Finish?
		...
		->
			Divorce the restful postgres concept from the Hub API
			Make one item a key priority
			Item: Allow reading from a postgres database and dynamic matching onto a struct
			!
			->
				This is the choice.
				First step: Clone the restful postgres repo


*/
