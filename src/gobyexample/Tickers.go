package main

import (
	"fmt"
	"time"
)

//tickers 是在某个时间重复执行，可以使用stop方法停止执行。<- ticker.C代表的是一次循环的结束
func main() {
	//定时执行某个活动。
	ticker := time.NewTicker(2 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at:", t)
			}
		}
	}()
	//上述 go func 代码重复执行，间隔2s

	time.Sleep(6 * time.Second)
	//手动停止tickers
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	/**输出结果
	tick at: 2020-09-11 14:51:51.431511 +0800 CST m=+2.004011863
	tick at: 2020-09-11 14:51:53.431683 +0800 CST m=+4.004123622
	tick at: 2020-09-11 14:51:55.431962 +0800 CST m=+6.004343159
	Ticker stopped
	*/
}
