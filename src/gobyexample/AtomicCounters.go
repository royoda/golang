package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
原子计数，多个协程对一个数操作，每个协程自增100次，50个协程即50*100 = 5000次的结果
*/
func main() {
	var ops uint64

	var wg sync.WaitGroup
	//var mutex sync.Mutex
	for i := 1; i <= 50; i++ {
		//老夫目前add里面的只能填1，填别的会报错，老夫暂时还不清楚原因
		wg.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				atomic.AddUint64(&ops, 1)
				/*mutex.Lock()
				ops++
				mutex.Unlock()*/
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)

}
