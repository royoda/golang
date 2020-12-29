package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//缓冲入口
	done <- true
}

//make(chan int, 1)和make(chan int) 的区别
/**
make(chan int) done <- true之后会阻塞，直到执行<-done才继续执行
make(chan int, 1) 使用<-done才会出现阻塞，等通道有值的情况才会继续往下走
*/
func main() {
	//定义通道，和缓冲值
	done := make(chan bool, 1)
	go worker(done)

	//缓冲出口，取出缓冲的值，如果缓冲没有赋值，即没有入口，则会报错
	<-done //如果删除该行代码，将会在程序执行worker方法前退出程序

	//make(chan int, 1)和make(chan int) 的区别
	/*var c = make(chan int, 1)
	var a string

	go func() {
		a = "hello world"
		<-c
	}()

	c <- 0
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(a)*/

	/*var c = make(chan int)
	var a string

	go func() {
		a = "hello world"
		<-c
	}()

	c <- 0
	fmt.Println(a)*/
}
