package main

import (
	"fmt"
	"time"
)

func workerPool(id int, jobs <-chan int, results chan<- int) {
	//遍历jobs通道，有数据就开始工作
	for j := range jobs {
		fmt.Println("worker:", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker:", id, "finished job", j)
		results <- j * 2
	}
}
func main() {
	const num = 5
	jobs := make(chan int, num)
	results := make(chan int, num)

	for w := 1; w <= 3; w++ {
		//创建3个工作协程
		go workerPool(w, jobs, results)
	}
	//添加工作，将数据放入通道中，协程会通过遍历通道获取到工作值就会开始工作。
	for i := 1; i <= num; i++ {
		jobs <- i
	}

	//关闭jobs通道,只用于发送数据，即send发送方，不能使用在recevied接受方。
	close(jobs)
	//释放通道数据
	for a := 1; a <= num; a++ {
		fmt.Println(<-results)
	}
}
