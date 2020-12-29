package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	//2s 后执行该方法
	fmt.Println("timer1 1 fired")

	ticker := time.NewTicker(2 * time.Second)
	fmt.Println("NewTicker")
	<-ticker.C

	timer2 := time.NewTimer(time.Second)
	//协程启动需要时间
	go func() {
		//自动停止计时器
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	//给予时间，使得协程可以启动
	time.Sleep(3 * time.Second)
	//手动停止计时器，当程序自定停止之后，stop2的值为false
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(3 * time.Second)
}
