package main

import (
	"fmt"
	"time"
)

//select 结合time.After 使让select超时拿不到结果而已打印，否则会打印报错信息
func main() {
	c1 := make(chan string, 1)

	go func() {
		//模拟超时情况
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println("res:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		//模拟操作
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c1:
		fmt.Println("res:", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
