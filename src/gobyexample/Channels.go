package main

import "fmt"

/*
	c3 := make(chan int, 1) 双向通道
	c3 := make(<-chan int, 1) 只读通道
	c2 := make(chan<- int, 1) 只写通道
*/
//channels 用于各个goroutine之间的通信
func ctest(msg chan int, num *int) {
	for i := 0; i < 3; i++ {
		*num++
	}
	msg <- *num

}

var num *int

func main() {
	//定义变量
	count1 := 0
	//给指针赋值
	num = &count1

	messages := make(chan int)
	//同步运算
	go ctest(messages, num)

	msg := <-messages
	fmt.Println(msg)
	//同步运算
	go func() {
		for i := 0; i < 3; i++ {
			*num++
		}
		messages <- *num
	}()

	msg1 := <-messages
	//<-messages 释放直接输出
	//最终结果
	fmt.Println(msg1)
}
