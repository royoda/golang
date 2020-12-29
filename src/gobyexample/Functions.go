package main

import "fmt"

//当您具有多个相同类型的连续参数时，可以忽略类似类型的参数的类型名称，直到声明该类型的最终参数为止。
//括号外的int是返回值类型
//a,b,c int
func hello(a int, b int) int {
	fmt.Println("hello")
	return a + b
}
func world() string {
	fmt.Println("world")
	return "hello"
}
func main() {
	res := hello(1, 2)
	fmt.Println(res)
	str := world()
	fmt.Println("str:", str)
}
