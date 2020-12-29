package main

import "fmt"

/*
多个返回值
*/
func rebackValues() (int, int) {

	return 0, 1
}
func main() {
	a, b := rebackValues()
	fmt.Println(a)
	fmt.Println(b)
	//如果只需要返回值的子集，请使用空白标识符_
	_, c := rebackValues()
	fmt.Println(c)
	d, _ := rebackValues()
	fmt.Println(d)
}
