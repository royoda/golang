package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//互斥锁
func main() {
	//定义map键值对
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	//创建100个协程,用与读操作
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				//具有原子性的自增操作，自增1
				atomic.AddUint64(&readOps, 1)
				//睡眠
				time.Sleep(time.Millisecond)
			}
		}()
	}
	//创建10个协程，用于写操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)   //随机返回一个整数，[0,5)
				val := rand.Intn(100) //随机返回一个整数，[0,100)
				//加锁操作
				mutex.Lock()
				state[key] = val
				//解锁操作
				mutex.Unlock()
				//具有原子性的自增操作，自增1
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second * 1)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOpsFinal:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
