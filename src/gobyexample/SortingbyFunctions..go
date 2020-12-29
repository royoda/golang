package main

import (
	"fmt"
	"sort"
)

//按功能排序
//自定义排序
//定义数组类型
type byLength []int

//sort 接口的内置方法 Len() int，Less(i, j int) bool，Swap(i, j int)
func (by byLength) Len() int {
	//返回数组长度
	return len(by)
}

//指定方式比较
func (by byLength) Less(i, j int) bool {
	return by[i] > by[j]
}

//交换元素
func (by byLength) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}
func main() {
	nums := []int{1, 3, 4, 2, 6, 7, 23}
	sort.Sort(byLength(nums))
	fmt.Println(nums)
}
