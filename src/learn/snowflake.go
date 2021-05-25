package main

import (
	"fmt"
	idworker "github.com/gitstliu/go-id-worker"
)

/**
    *@Description TODO
    *@Date 2021/5/24 2:59 下午
**/

func main() {
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newID, err := currWoker.NextId()
	if err == nil {
		fmt.Println(newID)
	}
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
