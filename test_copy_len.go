package main

import (
	"fmt"
)

func main() {
	/*const KeyPointItem = "PointItems:%s"
	key := fmt.Sprintf(KeyPointItem,
		time.Now().Add(-time.Hour*24).Format("2006-01-02"))*/
	//var a = 32 << 31
	// 测试copy方法
	var ar1 = []int{1,2,3}
	var ar2 = make([]int, len(ar1))
	arr := copy(ar2, ar1)
	fmt.Println(ar1)
	fmt.Println(ar2)
	fmt.Println(arr)

	fmt.Println(2 << 20)
	var obj = make([]byte,100)
	fmt.Println(len(obj))
	fmt.Println(cap(obj))
	fmt.Println(obj[:5])

}
