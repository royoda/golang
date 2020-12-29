package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func pingcount(pings chan<- int, msg int) {
	for i := 0; i < 3; i++ {
		msg++
	}
	pings <- msg
}

func pongcount(pings <-chan int, pongs chan<- int) {
	msg := <-pings
	for i := 0; i < 3; i++ {
		msg++
	}
	pongs <- msg
}

var count = 0

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	pingc := make(chan int, 1)
	pongc := make(chan int, 1)

	pingcount(pingc, count)
	pongcount(pingc, pongc)

	fmt.Println(<-pongc)
}
