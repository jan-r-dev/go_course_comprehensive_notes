package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)
	//receive
	receive(even, odd, quit)

	fmt.Println("exiting")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println(v, "is even.")
		case v := <-o:
			fmt.Println(v, "is odd.")
		case v := <-q:
			fmt.Println(v, "from the quit channel")
			return
		}
	}
}
