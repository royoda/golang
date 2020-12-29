package main

import "fmt"

func main() {
	c1 := make(chan string)
	singler := make(chan string)

	//case 没有获取到值，会立即执行default语句
	select {
	case msg := <-c1:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"
	//因为没有缓冲区，没有接受者，会直接运行default
	select {
	case c1 <- msg:
		fmt.Println("send message")
	default:
		fmt.Println("no send message")
	}

	//非阻塞操作
	select {
	case msg := <-c1:
		fmt.Println("received message", msg)
	case res := <-singler:
		fmt.Println("received message", res)
	default:
		fmt.Println("no activity")
	}

}
