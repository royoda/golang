package main

import "fmt"

/**
闭包
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//持续调用同一个函数的时候会一直递增。重新赋值函数的时候值会重新变成1，然后在继续开始递增
	nextInt := intSeq()
	//nextInt  0x1093150
	fmt.Println(nextInt)
	fmt.Println(&nextInt)
	fmt.Println(&nextInt)
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(nextInt())

}
