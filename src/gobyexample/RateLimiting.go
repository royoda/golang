package main

import (
	"fmt"
	"time"
)

func main() {

	request := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	//time.NewTicker 另一种形式，不过tick是用于通道，当通道数据为0时自动结束，不像newticker需要手动关闭
	//time.Tick 通道形式，每个一段时间返回通道内容
	limiter := time.Tick(200 * time.Millisecond)

	for req := range request {
		fmt.Println(<-limiter)
		fmt.Println("request:", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		//一开始会立即处理前面3个请求，因为后面使用了limiter，会间隔200毫秒后执行
		fmt.Println(<-burstyLimiter)
		fmt.Println("burstyLimiter", req, time.Now())
	}
}
