package main

import "fmt"

//数组的使用
func main() {
	var arr [5]int // 数组
	fmt.Println("arr", arr)
	var arr1 [5]int // 数组

	arr1 = arr //赋值之后修改arr的值，arr1的值不变。就是值复制，不是地址指向

	arr[4] = 100
	fmt.Println("arr", arr)
	fmt.Println("arr[4]", arr[4])
	fmt.Println("arr1", arr1)

	var arrs = [5]int{1, 2, 3, 4, 5}
	// arrs := [5]int {1,2,3,4,5}
	//var b = [...]int{1, 2, 3}
	for i := 0; i < len(arrs); i++ {
		fmt.Print(arrs[i], "\t")
	}
	fmt.Println()
	var twoD [2][3]int
	fmt.Println(twoD)
	//使用内置函数len获取数组的长度
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
