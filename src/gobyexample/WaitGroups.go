package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	程序运行后进入阻塞状态，
	等待所有goroutine运行完成，所有协程程序执行完毕，类似java语言的CountDownLatch程序计数的例子
	Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
*/
func workerGroup(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerGroup(i, &wg)
	}

	wg.Wait()
}
