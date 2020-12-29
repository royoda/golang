package main

import "fmt"

//for 是go语言中唯一的循环结构，不写循环条件默认死循环，直到遇到break，return等推出循环。
//也可以使用continue继续往下执行代码
func main() {
	i := 3
	//循环等价于 i:3;i<=3;i++,所以只会输出一次
	for i <= 3 {
		fmt.Println("i=", i)
		i = i + 1
	}

	fmt.Println("i=", i)
	//7，8，9
	for j := 7; j <= 9; j++ {
		fmt.Println("j=", j)
	}
	//死循环操作，终止条件时break
	for {
		fmt.Println("loop")
		if i == 6 {
			break
		}
		i++
	}
	//循环遍历，取出不可以整除2的数字
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
