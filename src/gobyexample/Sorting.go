package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	//排序字符串类型
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	//排序整数类型
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)
	//验证切片是否已经排序正常
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
