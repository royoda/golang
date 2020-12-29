package main

import "fmt"

/**
递归
*/
/*
	递归是压栈操作，压入 3 ，2 ，1，0等于0的时候返回1，在依次从栈中取出值做运算，1+1+2+3 = 7
*/
func recursion(i int) int {
	if i == 0 {
		return 1
	}
	return i + recursion(i-1)
}
func main() {
	res := recursion(5)
	fmt.Println(res)

}
