package main

import (
	"fmt"
	"time"
)

/**
协程：轻量级线程
*/
func f(from string) {
	for _, v := range from {
		fmt.Println(from, ":", string(v))
	}
}
func main() {
	f("hello")
	//主线程睡眠3秒
	time.Sleep(3000 * time.Millisecond)

	go f("world")
	go func(msg string) {
		for _, v := range msg {
			fmt.Println(msg, ":", string(v))
		}
	}("goroutines")
	//主线程睡眠3秒
	time.Sleep(3000 * time.Millisecond)

	fmt.Println("done")
}
