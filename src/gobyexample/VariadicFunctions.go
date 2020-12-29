package main

import "fmt"

/**
可变的参数
*/
func methon(nums ...int) {

	sum := 0
	for _, num := range nums {
		//sum =+ num 这相当于sum = num
		sum += num //sum = num + sum 达到循环添加数据的效果
	}
	fmt.Println(sum)
}
func main() {
	methon(1, 2, 3)
	methon(1, 2, 3, 4, 5, 6)

	nums := []int{1, 2, 3, 4, 5, 6}
	//使用数组当作参数的时候需要在后面用...补充
	methon(nums...)
}
