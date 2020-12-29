package main

import (
	"fmt"
	"time"
)

////select 会一直等待，直到等待到结果时将会进行输出。
//模拟例如阻止在并行goroutine中执行的RPC操作
func main() {
	//定义通道
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {

		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
