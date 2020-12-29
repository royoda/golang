package main

import "fmt"

//变量赋值
func main() {

	var a = "initial"
	fmt.Println(a) //initial

	var b, c int = 1, 2 //等于var b, c = 1, 2
	/**
	int yes
	double undefined
	float undefined
	boolean undefined
	string yes
	bool yes
	*/
	fmt.Println(b, c) //1,2
	var d = true
	fmt.Println(d) //true

	var e int
	fmt.Println(e) //0

	//该声明方式和var声明方式一样，赋值任意类型
	f := "apple"   //等价于var f string = "apple"
	fmt.Println(f) //apple

	var f1 = "apple"
	fmt.Println(f1)
}
