package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//信号
func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//注册
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //获取中断程序的命令，可由外部输入 例如：ctrl+c

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	//阻塞
	<-done
	fmt.Println("exiting")
}
