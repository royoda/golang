package main

import "fmt"

/**
range chan
*/
func main() {
	queue := make(chan string, 5)
	queue <- "h"
	queue <- "e"
	queue <- "l"
	queue <- "l"
	queue <- "o"

	close(queue)
	for v := range queue {
		fmt.Print(v)
	}
}
